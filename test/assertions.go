package test

import "testing"

func AssertWantErr(err error, wantErr, caller string, t *testing.T) bool {
	t.Helper()
	if err != nil {
		if wantErr != err.Error() {
			t.Errorf("%s error = %v, wantErr %q", caller, err, wantErr)
		}

		return true
	} else if wantErr != "" {
		t.Errorf("%s expected error %q, did not receive an error", caller, wantErr)
		return true
	}

	return false
}
