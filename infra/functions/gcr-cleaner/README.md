# GCR Cleaner

* GCR イメージ世代管理処理を行う関数を作成

## Payload &amp; Parameters

The payload is expected to be JSON with the following fields:

- `repo` - Full name of the repository to clean, in the format
  `gcr.io/project/repo`. This field is required.

- `keep` - If an integer is provided, it will always keep that minimum number
  of images. Note that it will not consider images inside the `grace` duration.

- `allow_tagged` - If set to true, will check all images including tagged.
  If unspecified, the default will only delete untagged images.

- `tag_filter` - Used for tags regexp definition to define pattern to clean,
  requires `allow_tagged` must be true. For example: use `-tag-filter "^dev.+$"`
  to limit cleaning only on the tags with begining with is `dev`. The default
  is no filtering. The regular expression is parsed according to the [Go regexp package syntax](https://golang.org/pkg/regexp/syntax/).

---

## 参考

* [GitHub - GCR Cleaner](https://github.com/sethvargo/gcr-cleaner)
* [GitHub - go-containerregistry](https://github.com/google/go-containerregistry)
