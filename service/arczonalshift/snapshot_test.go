// Code generated by smithy-go-codegen DO NOT EDIT.

//go:build snapshot

package arczonalshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	"io"
	"io/fs"
	"os"
	"testing"
)

const ssprefix = "snapshot"

type snapshotOK struct{}

func (snapshotOK) Error() string { return "error: success" }

func createp(path string) (*os.File, error) {
	if err := os.Mkdir(ssprefix, 0700); err != nil && !errors.Is(err, fs.ErrExist) {
		return nil, err
	}
	return os.Create(path)
}

func sspath(op string) string {
	return fmt.Sprintf("%s/api_op_%s.go.snap", ssprefix, op)
}

func updateSnapshot(stack *middleware.Stack, operation string) error {
	f, err := createp(sspath(operation))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(stack.String())); err != nil {
		return err
	}
	return snapshotOK{}
}

func testSnapshot(stack *middleware.Stack, operation string) error {
	f, err := os.Open(sspath(operation))
	if errors.Is(err, fs.ErrNotExist) {
		return snapshotOK{}
	}
	if err != nil {
		return err
	}
	defer f.Close()
	expected, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if actual := stack.String(); actual != string(expected) {
		return fmt.Errorf("%s != %s", expected, actual)
	}
	return snapshotOK{}
}
func TestCheckSnapshot_CancelPracticeRun(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CancelPracticeRun(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CancelPracticeRun")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CancelZonalShift(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CancelZonalShift(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CancelZonalShift")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreatePracticeRunConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreatePracticeRunConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreatePracticeRunConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeletePracticeRunConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeletePracticeRunConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeletePracticeRunConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetAutoshiftObserverNotificationStatus(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetAutoshiftObserverNotificationStatus(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetAutoshiftObserverNotificationStatus")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetManagedResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetManagedResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetManagedResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListAutoshifts(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListAutoshifts(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListAutoshifts")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListManagedResources(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListManagedResources(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListManagedResources")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListZonalShifts(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListZonalShifts(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListZonalShifts")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StartPracticeRun(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartPracticeRun(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StartPracticeRun")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StartZonalShift(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartZonalShift(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StartZonalShift")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateAutoshiftObserverNotificationStatus(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateAutoshiftObserverNotificationStatus(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateAutoshiftObserverNotificationStatus")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdatePracticeRunConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdatePracticeRunConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdatePracticeRunConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateZonalAutoshiftConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateZonalAutoshiftConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateZonalAutoshiftConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateZonalShift(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateZonalShift(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateZonalShift")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
func TestUpdateSnapshot_CancelPracticeRun(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CancelPracticeRun(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CancelPracticeRun")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CancelZonalShift(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CancelZonalShift(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CancelZonalShift")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreatePracticeRunConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreatePracticeRunConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreatePracticeRunConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeletePracticeRunConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeletePracticeRunConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeletePracticeRunConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetAutoshiftObserverNotificationStatus(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetAutoshiftObserverNotificationStatus(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetAutoshiftObserverNotificationStatus")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetManagedResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetManagedResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetManagedResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListAutoshifts(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListAutoshifts(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListAutoshifts")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListManagedResources(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListManagedResources(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListManagedResources")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListZonalShifts(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListZonalShifts(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListZonalShifts")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StartPracticeRun(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartPracticeRun(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StartPracticeRun")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StartZonalShift(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartZonalShift(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StartZonalShift")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateAutoshiftObserverNotificationStatus(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateAutoshiftObserverNotificationStatus(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateAutoshiftObserverNotificationStatus")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdatePracticeRunConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdatePracticeRunConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdatePracticeRunConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateZonalAutoshiftConfiguration(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateZonalAutoshiftConfiguration(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateZonalAutoshiftConfiguration")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateZonalShift(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateZonalShift(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateZonalShift")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
