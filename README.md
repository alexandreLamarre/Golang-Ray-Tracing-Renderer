# Golang-Ray-Tracing-Renderer
[![GoReportCard example](https://goreportcard.com/badge/github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer)](https://goreportcard.com/report/github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer)

A native golang 3D ray tracing renderer, that implements Ray intersection using the Phong Reflection Model and Constructive Solid Geometry, augmented with recursively defined Patterns, as well as Perlin/Simplex noise pattern algorihthms.

The native renderer runs on CPU cycles, so high resolution scenes may take some time to render (5-15mins)

Examples rendered using this code: 

<details>
  <summary> Basic refraction/ reflection (1400 x 1000)</summary>
  <img src = "https://github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/blob/main/examplesnative/basic%20reflect%20refract.png"/>
</details>
<details>
  <summary> Complex(ish) refraction/reflection (640 x 600) </summary>
   <img src = "https://github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/blob/main/examplesnative/complex%20reflect%20refract.png"/>
</details>

## Table of Contents
- [Native](#Native)
  - [Implementation Details](#Implementation-Details)
    - [Algebra](#Algebra)
    - [Canvas](#Canvas)
    - [Geometry](#Geometry)
    - [Camera](#Camera)
    - [Noise](#Noise)
    
-[BLAS gpu accelerated](#BLAS-gpu-accelerated)

### Native
[Back To Top](#)

#### Implementation Details
[Back To Top](#)

The native renderer is structured in such a way that each package handle one aspect of a 3D renderer or its own class of algorithms.

#### Algebra
[Back To Top](#)

The algebra package covers all the data structures, functions and methods used for algebraic manipulation of 64bit floating point numbers. It should only be used in the implementation of other packages and never in a standalone way.

#### Canvas
[Back To Top](#)

The canvas package covers all the data structures, functions and methods relating to representing pixels, exporting pixels to files, colors, patterns, Point lights and materials.

#### Geometry
[Back To Top](#)

The geometry package covers all the data structures, functions and methods relating to the World Space and the Shape interface (and by extension all the basic shapes provided in the package)

#### Camera
[Back To Top](#)

The Camera package covers the camera data structure, and its functions and methods, to contruct a geometric pipeline that transforms the world space coordinates into their view space coordinates and pixel space coordinates.

#### Noise
[Back To Top](#)

The noise package covers Perlin/Simplex noise algorithms. Unittested using the original java implementation outputs and seeds and comparing it to the outputs of the Golang traslated algorithms using the same seed.

