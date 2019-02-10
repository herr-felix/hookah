
export default
class BuildAPI {
  static async getOutputById(buildId: string): Promise<string> {
    return  `00:00:00 apk add musl-dev go
            00:00:00 fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/main/x86_64/APKINDEX.tar.gz
            00:00:00 fetch http://dl-cdn.alpinelinux.org/alpine/v3.8/community/x86_64/APKINDEX.tar.gz
            00:00:01 (1/13) Installing binutils (2.30-r5)
            00:00:01 (2/13) Installing gmp (6.1.2-r1)
            00:00:01 (3/13) Installing isl (0.18-r0)
            00:00:01 (4/13) Installing libgomp (6.4.0-r9)
            00:00:01 (5/13) Installing libatomic (6.4.0-r9)
            00:00:02 (6/13) Installing pkgconf (1.5.3-r0)
            00:00:02 (7/13) Installing libgcc (6.4.0-r9)
            00:00:02 (8/13) Installing mpfr3 (3.1.5-r1)
            00:00:02 (9/13) Installing mpc1 (1.0.3-r1)
            00:00:02 (10/13) Installing libstdc++ (6.4.0-r9)
            00:00:02 (11/13) Installing gcc (6.4.0-r9)
            00:00:09 (12/13) Installing go (1.10.8-r0)
            00:00:33 (13/13) Installing musl-dev (1.1.19-r10)
            00:00:34 Executing busybox-1.28.4-r3.trigger
            00:00:34 OK: 456 MiB in 40 packages
            00:00:34 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
            00:00:38 echo "All done"
            00:00:38 All done`
  }
}