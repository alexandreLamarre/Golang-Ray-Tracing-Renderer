# Golang-Ray-Tracing-Renderer

[![GoReportCard example](https://goreportcard.com/badge/github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer)](https://goreportcard.com/report/github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer)


A native (no external libraries) golang 3D ray tracing renderer, that implements Ray intersection via the [Phong-Reflection Model](https://en.wikipedia.org/wiki/Phong_reflection_model), Recursive Shape grouping with [AABB optimizations](https://en.wikipedia.org/wiki/Bounding_volume) and [Constructive Solid Geometry](https://en.wikipedia.org/wiki/Constructive_solid_geometry), augmented with recursively defined Patterns, [Perlin/Simplex noise algorihthms](https://en.wikipedia.org/wiki/Perlin_noise) (for patterns or triangulated objects), efficient convex-hull/half-edge [Delaunay Triangulation](https://en.wikipedia.org/wiki/Delaunay_triangulation) implementation and a [LindenMayer-System](https://en.wikipedia.org/wiki/L-system) fractal interface (comes with a [mandelbulb](https://en.wikipedia.org/wiki/Mandelbulb) implementation).

The native renderer runs on CPU cycles, so high resolution scenes may take some time to render (5-15mins). Pre-rendering optimizations on .obj files relies on faces that are ordered by position: i.e. that the closest face to the previous face comes directly after it in the .obj definition.

It renders and writes to .ppm files which can be opened natively on MacOS with preview but require some additional software on Windows/Linux. I recommend using [GIMP](https://www.gimp.org/) to open these files because it is a well-maintained open source Image Manipulation Program.

**Examples rendered using this code**:

<details>
  <summary> Basic refraction/ reflection (1400 x 1000)</summary>
  <img src = "https://github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/blob/main/examplesnative/basic%20reflect%20refract.png"/>
</details>
<details>
  <summary> Complex(ish) refraction/reflection (1400 x 1000) </summary>
   <img src = "https://github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/blob/main/examplesnative/complex%20reflect%20refract.png"/>
</details>

Some notable improvements to be made: 
- making a fast algebra library that binds algebraic manipulations and their data structures to the gpu.
- replace `getBounds()` method for Groups when Adding a `Shape` to a `Group` to a function that finds a fast AABB union for the `Bounds` struct

## Table of Contents
- [Usage](#Usage)
- [Implementation Details](#Implementation-Details)
  - [Algebra](#Algebra)
  - [Canvas](#Canvas)
  - [Geometry](#Geometry)
  - [Camera](#Camera)
  - [Noise](#Noise)
    
-[BLAS gpu accelerated](#BLAS-gpu-accelerated)

### Usage
[Back To Top](#)

To use this repository, you must first fork it and the clone it to your desktop.

You can then open a terminal window and cd to its root directory and use the following CLI commands to run it on some preconfigured actions:

#### Parsing a .obj file
[Back To Top](#)

You can run this command to parse a .obj file and render/write it to a .ppm file in `pkg/examples/ directory`:

`go run main.go -p -parsefile=<string:filepath/filename> -name=<string>`

- The parsefile flag accepts a filepath from the root directory and a filename to open and parse.
- The name flag accepts a string that specificies the name of the file saved to `./pkg/examples/<string:name>.ppm`

### Implementation Details
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

