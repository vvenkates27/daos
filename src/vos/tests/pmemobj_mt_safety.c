/**
 * (C) Copyright 2020 Intel Corporation.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
 * The Government's rights to use, modify, reproduce, release, perform, display,
 * or disclose this software are subject to the terms of the Apache License as
 * provided in Contract No. B609815.
 * Any reproduction of computer software, computer software documentation, or
 * portions thereof marked with this legend must also reproduce the markings.
 */
/**
 * Utility to test the pmemobj pool locking/thread safety for pmemobj pool
 * operations
 * Author: Li Wei G  <wei.g.li@intel.com>
 */

#define D_LOGFAC		DD_FAC(tests)

#include <fcntl.h>
#include <pthread.h>
#include <assert.h>
#include <sys/types.h>
#include <stdio.h>
#include <sys/syscall.h>
#include <stdlib.h>
#include <libpmemobj.h>
#include <unistd.h>
#include <stdarg.h>

static pthread_barrier_t barrier;

static inline pid_t
gettid(void)
{
	return syscall(SYS_gettid);
}

static void *
func(void *arg)
{
	char	       *prefix = arg;
	char		path[1024];
	PMEMobjpool    *pool;
	int		n = 10;
	int		fd;
	int		i;
	int		rc;

	snprintf(path, sizeof path, "%s-%ld", prefix, (long)gettid());

	printf("%ld: creating %s\n", (long)gettid(), path);
	fd = open(path, O_CREAT|O_RDWR, 0600);
	assert(fd >= 0);
	rc = posix_fallocate(fd, 0, (1 << 24) /* 16 MB */);
	assert(rc == 0);
	pool = pmemobj_create(path, NULL /* layout */, 0 /* poolsize */,
			      0644 /* mode */);
	if (pool == NULL) {
		printf("%ld: failed to create %s: %d\n", (long)gettid(), path,
			errno);
		goto out;
	}
	pmemobj_close(pool);

	printf("%ld: waiting for barrier\n", (long)gettid());
	rc = pthread_barrier_wait(&barrier);
	assert(rc == 0 || rc == PTHREAD_BARRIER_SERIAL_THREAD);

	for (i = 0; i < n; i++) {
		pool = pmemobj_open(path, NULL /* layout */);
		if (pool == NULL) {
			printf("%ld: failed to open %s: %d\n", (long)gettid(),
				path, errno);
			goto out_file;
		}
		pmemobj_close(pool);
	}

out_file:
	printf("%ld: destroying %s\n", (long)gettid(), path);
	rc = unlink(path);
	assert(rc == 0);

out:
	return NULL;
}

int
main(int argc, char *argv[])
{
	char	       *path = "/mnt/daos/pmemobj_mt_safety";
	int		nthreads = 32;
	pthread_t      *threads;
	int		i;
	int		rc;

	threads = malloc(sizeof(*threads) * nthreads);
	assert(threads != NULL);

	rc = pthread_barrier_init(&barrier, NULL /* attr */, nthreads);
	assert(rc == 0);

	for (i = 0; i < nthreads; i++) {
		rc = pthread_create(&threads[i], NULL /* attr */, func, path);
		assert(rc == 0);
	}

	for (i = 0; i < nthreads; i++) {
		rc = pthread_join(threads[i], NULL /* retval */);
		assert(rc == 0);
	}

	pthread_barrier_destroy(&barrier);
	free(threads);
	return 0;
}
