package vec3

type Ray struct {
	origin Vec3
	dir    Vec3
}

func (r *Ray) GetOrigin() Vec3 {
	return r.origin
}

func (r *Ray) GetDir() Vec3 {
	return r.dir
}

func (r *Ray) SetOrigin(v Vec3) {
	r.origin = v
}

func (r *Ray) SetDir(v Vec3) {
	r.dir = v
}

func (r *Ray) PointAt(t float64) Vec3 {
	//Computes new point P = O + dir(t)
	multDir := r.dir.MultiplybyFloat(t)
	newPoint := r.origin.Add(multDir)
	return *newPoint
}
