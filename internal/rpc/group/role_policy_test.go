package group

import "testing"

func TestCanUpdateMemberRole(t *testing.T) {
	const (
		groupOwner      int32 = 100
		groupAdmin      int32 = 60
		ordinaryMember int32 = 20
	)

	tests := []struct {
		name          string
		operatorRole  int32
		changingRole  bool
		want          bool
	}{
		{name: "owner may change member role", operatorRole: groupOwner, changingRole: true, want: true},
		{name: "admin may not change member role", operatorRole: groupAdmin, changingRole: true, want: false},
		{name: "ordinary member may not change member role", operatorRole: ordinaryMember, changingRole: true, want: false},
		{name: "owner may update non-role member info", operatorRole: groupOwner, changingRole: false, want: true},
		{name: "admin may update non-role member info", operatorRole: groupAdmin, changingRole: false, want: true},
		{name: "ordinary member may update non-role member info", operatorRole: ordinaryMember, changingRole: false, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canUpdateMemberRole(tt.operatorRole, tt.changingRole); got != tt.want {
				t.Errorf("canUpdateMemberRole(%d, %t) = %t, want %t", tt.operatorRole, tt.changingRole, got, tt.want)
			}
		})
	}
}
