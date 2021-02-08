# Scumbag · 渣男

**Building private media cloud for Online education, entertainment and streaming**

## Features Supported

### Uploading

- [x] Standard upload
- [ ] Chunked upload
- [ ] File Transfer Protocol

### Object Storage

- [x] Local
- [ ] Minio
- [ ] ceph
- [ ] S3

### Transcoding
- [ ] documents
- [ ] videos
- [ ] audios

## Feature

### Uploading

```http request
PUT /upload HTTP/1.1
Content-Name: your-file-origin-name
Digest: SHA-256=sha256-of-uploading-file
Host: localhost:4200
Content-Length: 597
your file content
```