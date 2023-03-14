package car

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		// if call fatal the test will not be run through
		t.Error("got errors", err)
	}

	if c == nil {
		t.Error("car should not be nil")
	}
}

func TestNewWithAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = New("foo", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "foo", c.Name)
}

func TestCar_SetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    *Car
		args args
		want string
	}{
		{
			name: "no input name",
			c: &Car{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "",
			},
			want: "foo",
		}, {
			name: "input name",
			c: &Car{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "bar",
			},
			want: "bar",
		}, {
			name: "init name with no args",
			c: &Car{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "",
			},
			want: "foo",
		},
	}
	for _, tt := range tests {
		tt := tt // new the variable when run parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // Will parallel run test use go routine

			if got := tt.c.SetName(tt.args.name); got != tt.want {
				t.Errorf("Car.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
