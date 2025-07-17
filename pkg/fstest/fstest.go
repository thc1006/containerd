
// ChownFunc is used by Chown() to change ownership.
var ChownFunc = os.Chown

// Chown changes uid/gid of the given path and retries on transient failures.
func Chown(path string, uid, gid int) Applier {
	return applyFn(func(root string) error {
		abs := filepath.Join(root, path)
		return Retry(context.Background(), 5, 10*time.Millisecond, func() error {
			return ChownFunc(abs, uid, gid)
		})
	})
}
