#include <immintrin.h>
#include "reader.h"

/*
 * Create a new reader.
 */
reader_t *new_reader(const void *src, size_t src_size)
{
    reader_t *r = malloc(sizeof *r);
    r->src = src;
    r->src_size = src_size;
    r->off = 0;
    return r;
}

/*
 * Copy a pattern into a buffer with an offset into the pattern.
 *
 * This can be used to repeatedly fill a buffer with a portion of data, creating a stream of the pattern.
 *
 * Pseudocode example:
 *
 * \code
 * pat <- "foo"
 * buf <- "__"
 *
 * off <- copy_with_offset(buf, pat, 2, 3, 0)
 *
 * off: 2
 * buf: "fo"
 *
 * off <- copy_with_offset(buf, pat, 2, 3, off)
 *
 * off: 1
 * buf: "of"
 *
 * off <- copy_with_offset(buf, pat, 2, 3, off)
 *
 * off: 0
 * buf: "oo"
 *
 * off <- copy_with_offset(buf, pat, 2, 3, off)
 *
 * off: 2
 * buf: "fo"
 * \endcode
 *
 * @param r        the reader_t holding any state
 * @param dst      pointer to the start of the buffer
 * @param dst_size number of bytes in the buffer. Must be a multiple of sizeof(__uint128_t)!
 *
 * @return dst_size to resemble with Go's io.Reader
 */
size_t _read(reader_t *r, void *dst, size_t dst_size)
{
    __uint128_t *d = dst;
    const __uint128_t *s = r->src;

    size_t maxOff = r->src_size / sizeof(__uint128_t);
    for (size_t i = 0; i < dst_size / sizeof(__uint128_t); i++)
    {
        d[i] = s[r->off];

        r->off++;
        if (r->off == maxOff)
        {
            r->off = 0;
        }
    }

    return dst_size;
}
