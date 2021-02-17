package examples

import "github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"

type projectile struct {
	position *algebra.Vector
	velocity *algebra.Vector
}

type environment struct {
	gravity *algebra.Vector
	wind    *algebra.Vector
}

func tick(env *environment, proj *projectile) error {
	newPos, err := proj.position.Add(proj.velocity)
	if err != nil {
		return err
	}
	proj.position = newPos
	newVelocity, err := proj.velocity.Add(env.gravity)
	if err != nil {
		return err
	}
	newVelocity, err = newVelocity.Add(env.wind)
	if err != nil {
		return err
	}
	proj.velocity = newVelocity
	return nil
}
