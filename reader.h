#include <stddef.h>

typedef struct {
    const void *src;
    size_t src_size;
    size_t off;
} reader_t;

reader_t* new_reader(const void *src, size_t src_size);

size_t _read(reader_t *r, void *dst, size_t dst_size);
