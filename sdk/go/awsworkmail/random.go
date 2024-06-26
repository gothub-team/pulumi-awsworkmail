// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsworkmail

import (
	"context"
	"reflect"

	"errors"
	"github.com/gothub-team/pulumi-awsworkmail/sdk/go/awsworkmail/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type Random struct {
	pulumi.CustomResourceState

	Length pulumix.Output[int]    `pulumi:"length"`
	Result pulumix.Output[string] `pulumi:"result"`
}

// NewRandom registers a new resource with the given unique name, arguments, and options.
func NewRandom(ctx *pulumi.Context,
	name string, args *RandomArgs, opts ...pulumi.ResourceOption) (*Random, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Random
	err := ctx.RegisterResource("awsworkmail:index:Random", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandom gets an existing Random resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomState, opts ...pulumi.ResourceOption) (*Random, error) {
	var resource Random
	err := ctx.ReadResource("awsworkmail:index:Random", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Random resources.
type randomState struct {
}

type RandomState struct {
}

func (RandomState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomState)(nil)).Elem()
}

type randomArgs struct {
	Length int `pulumi:"length"`
}

// The set of arguments for constructing a Random resource.
type RandomArgs struct {
	Length pulumix.Input[int]
}

func (RandomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomArgs)(nil)).Elem()
}

type RandomOutput struct{ *pulumi.OutputState }

func (RandomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Random)(nil)).Elem()
}

func (o RandomOutput) ToRandomOutput() RandomOutput {
	return o
}

func (o RandomOutput) ToRandomOutputWithContext(ctx context.Context) RandomOutput {
	return o
}

func (o RandomOutput) ToOutput(ctx context.Context) pulumix.Output[Random] {
	return pulumix.Output[Random]{
		OutputState: o.OutputState,
	}
}

func (o RandomOutput) Length() pulumix.Output[int] {
	value := pulumix.Apply[Random](o, func(v Random) pulumix.Output[int] { return v.Length })
	return pulumix.Flatten[int, pulumix.Output[int]](value)
}

func (o RandomOutput) Result() pulumix.Output[string] {
	value := pulumix.Apply[Random](o, func(v Random) pulumix.Output[string] { return v.Result })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(RandomOutput{})
}
