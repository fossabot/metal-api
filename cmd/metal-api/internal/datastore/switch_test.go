package datastore

import (
	"reflect"
	"testing"

	"git.f-i-ts.de/cloud-native/metal/metal-api/cmd/metal-api/internal/metal"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

func TestRethinkStore_FindSwitch(t *testing.T) {

	// mock the DB
	//ds, mock := initMockDB()
	//mock.On(r.DB("mockdb").Table("switch").Get("2")).Return(switch2, nil)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		want    *metal.Switch
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		/*
			{
				name: "TestRethinkStore_FindSwitch Test 1",
				rs:   ds,
				args: args{
					id: "2",
				},
				want:    &switch2,
				wantErr: false,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rs.FindSwitch(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.FindSwitch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RethinkStore.FindSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRethinkStore_findSwitchByRack(t *testing.T) {

	// mock the DB
	//[{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}]
	//[{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}]
	/*
		ds, mock := initMockDB()
		mock.On(r.DB("mockdb").Table("switch").Filter(func(var_1 r.Term) r.Term { return var_1.Field("rackid").Eq("rack1") })).Return([]metal.Switch{
			switch2,
		}, nil)
	*/
	type args struct {
		rackid string
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		want    []metal.Switch
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		/*
			{
				name: "TestRethinkStore_findSwitchByRack Test 1",
				rs:   ds,
				args: args{
					rackid: "rack1",
				},
				want: []metal.Switch{
					switch2,
				},
				wantErr: false,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rs.findSwitchByRack(tt.args.rackid)
			if (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.findSwitchByRack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RethinkStore.findSwitchByRack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRethinkStore_ListSwitches(t *testing.T) {

	// [{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}]
	// [{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}]

	// mock the DBs
	/*
		ds, mock := initMockDB()
		mock.On(r.DB("mockdb").Table("switch")).Return([]metal.Switch{
			switch1, switch2,
		}, nil)
		ds2, mock2 := initMockDB()
		mock2.On(r.DB("mockdb").Table("switch")).Return([]metal.Switch{
			switch2,
		}, nil)
	*/
	tests := []struct {
		name    string
		rs      *RethinkStore
		want    []metal.Switch
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		/*
			{
				name: "TestRethinkStore_ListSwitches Test 1",
				rs:   ds,
				want: []metal.Switch{
					switch1, switch2,
				},
				wantErr: false,
			},
			{
				name: "TestRethinkStore_ListSwitches Test 2",
				rs:   ds2,
				want: []metal.Switch{
					switch2,
				},
				wantErr: false,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rs.ListSwitches()
			if (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.ListSwitches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RethinkStore.ListSwitches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRethinkStore_CreateSwitch(t *testing.T) {

	// Diff. to Size and Site: They return nil, this switch returns the created switch. ==> is it wanted like this??

	// mock the DB
	ds, mock := initMockDB()
	mock.On(r.DB("mockdb").Table("switch").Insert(switch1)).Return(emptyResult, nil)

	type args struct {
		s *metal.Switch
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		want    *metal.Switch
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		{
			name: "TestRethinkStore_CreateSwitch Test 1",
			rs:   ds,
			args: args{
				s: &switch1,
			},
			want:    &switch1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rs.CreateSwitch(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.CreateSwitch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RethinkStore.CreateSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRethinkStore_DeleteSwitch(t *testing.T) {
	/*
		&{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}
		&{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}

		// mock the DBs
		ds, mock := initMockDB()
		mock.On(r.DB("mockdb").Table("switch").Get("switch1")).Return(switch1, nil)
		mock.On(r.DB("mockdb").Table("switch").Get("switch1").Delete()).Return(emptyResult, nil)
		mock.On(r.DB("mockdb").Table("switch").Get("switch2")).Return(switch2, nil)
		mock.On(r.DB("mockdb").Table("switch").Get("switch2").Delete()).Return(emptyResult, nil)
		mock.On(r.DB("mockdb").Table("switch").Get("switch3")).Return(emptyResult, nil)
		mock.On(r.DB("mockdb").Table("switch").Get("switch3").Delete()).Return(emptyResult, r.ErrEmptyResult)
	*/
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		want    *metal.Switch
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		/*
			{
				name: "TestRethinkStore_DeleteSwitch Test 1",
				rs:   ds,
				args: args{
					id: "switch1",
				},
				want:    &switch1,
				wantErr: false,
			},
			{
				name: "TestRethinkStore_DeleteSwitch Test 2",
				rs:   ds,
				args: args{
					id: "switch2",
				},
				want:    &switch2,
				wantErr: false,
			},
			{
				name: "TestRethinkStore_DeleteSwitch Test 3",
				rs:   ds,
				args: args{
					id: "switch3",
				},
				want:    nil,
				wantErr: true,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rs.DeleteSwitch(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.DeleteSwitch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RethinkStore.DeleteSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRethinkStore_UpdateSwitch(t *testing.T) {

	// mock the DB
	ds, mock := initMockDB()
	mock.On(r.DB("mockdb").Table("switch").Get("switch2").Replace(func(row r.Term) r.Term {
		return r.Branch(row.Field("changed").Eq(r.Expr(switch2.Changed)), switch3, r.Error("the switch was changed from another, please retry"))
	})).Return(emptyResult, nil)
	mock.On(r.DB("mockdb").Table("switch").Get("switch3").Replace(func(row r.Term) r.Term {
		return r.Branch(row.Field("changed").Eq(r.Expr(switch3.Changed)), switch2, r.Error("the switch was changed from another, please retry"))
	})).Return(emptyResult, nil)

	type args struct {
		oldSwitch *metal.Switch
		newSwitch *metal.Switch
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		{
			name: "TestRethinkStore_UpdateSwitch Test 1",
			rs:   ds,
			args: args{
				&switch2, &switch3,
			},
			wantErr: false,
		},
		{
			name: "TestRethinkStore_UpdateSwitch Test 2",
			rs:   ds,
			args: args{
				&switch3, &switch2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rs.UpdateSwitch(tt.args.oldSwitch, tt.args.newSwitch); (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.UpdateSwitch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRethinkStore_UpdateSwitchConnections(t *testing.T) {

	// mock the DB
	ds, mock := initMockDB()
	mock.On(r.DB("mockdb").Table("switch").Get("switch2").Replace(func(row r.Term) r.Term {
		return r.Branch(row.Field("changed").Eq(r.Expr(switch2.Changed)), switch2,
			r.Error("the switch was changed from another, please retry"))
	})).Return(emptyResult, nil)

	mock.On(r.DB("mockdb").Table("switch").Get("switch3").Replace(func(row r.Term) r.Term {
		return r.Branch(row.Field("changed").Eq(r.Expr(switch3.Changed)), switch3, r.Error("the switch was changed from another, please retry"))
	})).Return(emptyResult, nil)
	mock.On(r.DB("mockdb").Table("switch").Filter(func(var_1 r.Term) r.Term { return var_1.Field("rackid").Eq("rack1") })).Return([]metal.Switch{
		switch2,
	}, nil)
	mock.On(r.DB("mockdb").Table("switch")).Return([]metal.Switch{
		switch2, switch3,
	}, nil)

	type args struct {
		dev *metal.Device
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		{
			name: "TestRethinkStore_UpdateSwitchConnections Test 1",
			rs:   ds,
			args: args{
				&d1,
			},
			wantErr: false,
		},
		{
			name: "TestRethinkStore_UpdateSwitchConnections Test 2",
			rs:   ds,
			args: args{
				&d2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rs.UpdateSwitchConnections(tt.args.dev); (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.UpdateSwitchConnections() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRethinkStore_findSwithcByMac(t *testing.T) {
	/*
		[{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1} {{switch3   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}], want
		[{{switch2   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1} {{switch3   0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} [] [] map[] 1 rack1}]
	*/
	/*
		// mock the DB
		ds, mock := initMockDB()
		mock.On(r.DB("mockdb").Table("switch").Get("switch2").Replace(func(row r.Term) r.Term {
			return r.Branch(row.Field("changed").Eq(r.Expr(switch2.Changed)), switch2,
				r.Error("the switch was changed from another, please retry"))
		})).Return(emptyResult, nil)

		mock.On(r.DB("mockdb").Table("switch").Get("switch3").Replace(func(row r.Term) r.Term {
			return r.Branch(row.Field("changed").Eq(r.Expr(switch3.Changed)), switch3, r.Error("the switch was changed from another, please retry"))
		})).Return(emptyResult, nil)
		mock.On(r.DB("mockdb").Table("switch").Filter(func(var_1 r.Term) r.Term { return var_1.Field("rackid").Eq("rack1") })).Return([]metal.Switch{
			switch2,
		}, nil)
		mock.On(r.DB("mockdb").Table("switch")).Return([]metal.Switch{
			switch2, switch3,
		}, nil)

		var macs = []metal.Nic{
			nic1, nic2,
		}
		var searchmacs []string
		for _, m := range macs {
			searchmacs = append(searchmacs, string(m.MacAddress))
		}
		macexpr := r.Expr(searchmacs)

		mock.On(r.DB("mockdb").Table("switch").Filter(func(row r.Term) r.Term {
			return macexpr.SetIntersection(row.Field("network_interfaces").Field("macAddress")).Count().Gt(1)
		})).Return([]metal.Switch{
			switch2, switch3,
		}, nil)
	*/
	type args struct {
		macs []metal.Nic
	}
	tests := []struct {
		name    string
		rs      *RethinkStore
		args    args
		want    []metal.Switch
		wantErr bool
	}{
		// Test Data Array / Test Cases:
		/*
			{
				name: "TestRethinkStore_findSwithcByMac Test 1",
				rs:   ds,
				args: args{
					macs: []metal.Nic{
						nic1, nic2,
					},
				},
				want: []metal.Switch{
					switch2, switch3,
				},
				wantErr: false,
			},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.rs.findSwithcByMac(tt.args.macs)
			if (err != nil) != tt.wantErr {
				t.Errorf("RethinkStore.findSwithcByMac() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RethinkStore.findSwithcByMac() = %v, want %v", got, tt.want)
			}
		})
	}
}