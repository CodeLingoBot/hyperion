package light

import "context"

//GenericLight is for testing
type GenericLight struct {
	Name string `json:"name" yaml:"name"`
}

//GetName returns the light's name.
func (gl *GenericLight) GetName() string {
	return gl.Name
}

//GetType returns the type of light.
func (gl *GenericLight) GetType() string {
	return TypeGeneric
}

//GetID returns a fake light id
func (gl *GenericLight) GetID() string {
	return "generic-n/a"
}

//SetState updates the light's state.
func (gl *GenericLight) SetState(ctx context.Context, m Manager, target TargetState) {
	m.SetState(gl.Name, target.ToState())

}
