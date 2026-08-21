# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Tested | Latest | Activity |
|-----|--------|--------|----------|
| [Ark](https://github.com/mlange-42/ark) | v0.8.3 | ![GitHub Tag](https://img.shields.io/github/v/tag/mlange-42/ark?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/mlange-42/ark?label=date) | ![Last commit](https://img.shields.io/github/last-commit/mlange-42/ark) |
| [Donburi](https://github.com/yohamta0/donburi-ecs) | v1.15.7 | ![GitHub Tag](https://img.shields.io/github/v/tag/yohamta0/donburi-ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/yohamta0/donburi-ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/yohamta0/donburi-ecs) |
| [go‑gameengine‑ecs](https://github.com/marioolofo/go-gameengine-ecs) | v0.9.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/marioolofo/go-gameengine-ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/marioolofo/go-gameengine-ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/marioolofo/go-gameengine-ecs) |
| [GOKe](https://github.com/kjkrol/goke) | v1.2.6 | ![GitHub Tag](https://img.shields.io/github/v/tag/kjkrol/goke?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/kjkrol/goke?label=date) | ![Last commit](https://img.shields.io/github/last-commit/kjkrol/goke) |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | v0.0.3 | ![GitHub Tag](https://img.shields.io/github/v/tag/unitoftime/ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/unitoftime/ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/unitoftime/ecs) |
| [Volt](https://github.com/akmonengine/volt) | v1.7.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/akmonengine/volt?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/akmonengine/volt?label=date) | ![Last commit](https://img.shields.io/github/last-commit/akmonengine/volt) |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.
See the section on [Contributing](#contributing) for details.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Features

| ECS | Type-safe API | ID-based API | Relations | Events<sup>[1]</sup> | Batches<sup>[2]</sup> | Command buffer |
|-----|:-------------:|:------------:|:---------:|:-------:|:-------:|:--------------:|
| [Ark](https://github.com/mlange-42/ark) | ✅ | ✅ | ✅ | ✅ | ✅ | ❌ |
| [Donburi](https://github.com/yohamta0/donburi-ecs) | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |
| [go‑gameengine‑ecs](https://github.com/marioolofo/go-gameengine-ecs) | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| [GOKe](https://github.com/kjkrol/goke) | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |
| [Volt](https://github.com/akmonengine/volt) | ✅ | ❌ | ❌ | ✅ | ❌ | ❌ |

[1] ECS lifecycle events, allowing to react to entity creation, component addition, ...  
[2] Faster batch operations for entity creation etc.

## Benchmarks

Last run: Fri, 21 Aug 2026 20:09:47 UTC  
CPU: AMD EPYC 9V74 80-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

Note that the Y axis has logarithmic scale in all plots.
So doubled bar or line height is not doubled time!

All components used in the benchmarks have two `float64` fields.
The initial capacity of the world is set to 1024 where this is supported.

### Query

`N` entities with components `Position` and `Velocity`.
10x `N` entities with components `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](docs/results/query2comp.svg)

| N | Ark | Ark (tables) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 67.85ns | 71.51ns | 67.01ns | 55.23ns | 20.85ns | 16.79ns | 82.51ns |
| 4 | 18.77ns | 18.39ns | 30.97ns | 17.37ns | 5.74ns | 6.53ns | 21.59ns |
| 16 | 6.68ns | 5.29ns | 21.98ns | 7.74ns | 2.01ns | 3.93ns | 6.43ns |
| 64 | 3.91ns | 1.87ns | 20.64ns | 6.37ns | 1.05ns | 3.28ns | 2.70ns |
| 256 | 3.12ns | 1.05ns | 19.83ns | 6.10ns | 0.81ns | 3.20ns | 1.74ns |
| 1k | 2.16ns | 0.91ns | 19.95ns | 6.04ns | 0.80ns | 3.25ns | 1.54ns |
| 16k | 2.52ns | 0.79ns | 22.44ns | 6.00ns | 0.85ns | 3.15ns | 1.43ns |
| 256k | 2.89ns | 0.83ns | 23.86ns | 6.06ns | 0.91ns | 3.22ns | 1.43ns |
| 1M | 2.95ns | 0.92ns | 27.77ns | 6.04ns | 1.60ns | 3.20ns | 1.43ns |


> **Note:** Donburi, GOKe, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](docs/results/query32arch.svg)

| N | Ark | Ark (tables) | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 67.90ns | 72.62ns | 68.32ns | 19.50ns | 55.86ns | 16.28ns | 87.02ns |
| 4 | 35.20ns | 41.31ns | 34.41ns | 15.41ns | 22.40ns | 15.54ns | 72.24ns |
| 16 | 24.81ns | 33.21ns | 28.64ns | 14.39ns | 14.23ns | 22.86ns | 64.76ns |
| 64 | 13.93ns | 17.83ns | 23.41ns | 9.05ns | 8.96ns | 13.37ns | 33.02ns |
| 256 | 5.77ns | 5.74ns | 21.51ns | 3.70ns | 6.74ns | 6.10ns | 9.95ns |
| 1k | 3.17ns | 2.19ns | 21.25ns | 1.66ns | 6.21ns | 3.91ns | 3.58ns |
| 16k | 2.27ns | 1.00ns | 24.23ns | 0.96ns | 6.04ns | 3.27ns | 1.77ns |
| 256k | 2.16ns | 0.87ns | 73.70ns | 0.97ns | 6.05ns | 3.19ns | 1.55ns |
| 1M | 2.19ns | 1.02ns | 96.25ns | 1.88ns | 6.04ns | 3.24ns | 1.55ns |


> **Note:** Donburi, GOKe, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](docs/results/query256arch.svg)

| N | Ark | Ark (tables) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 65.95ns | 70.26ns | 66.57ns | 70.05ns | 19.55ns | 16.24ns | 105.89ns |
| 4 | 19.17ns | 18.19ns | 30.67ns | 30.10ns | 5.74ns | 9.43ns | 50.84ns |
| 16 | 6.79ns | 5.16ns | 21.41ns | 21.11ns | 2.06ns | 4.69ns | 39.53ns |
| 64 | 3.82ns | 1.84ns | 19.51ns | 19.88ns | 1.07ns | 3.45ns | 39.70ns |
| 256 | 2.35ns | 1.03ns | 19.89ns | 9.76ns | 0.81ns | 3.19ns | 10.86ns |
| 1k | 2.24ns | 0.95ns | 19.67ns | 7.00ns | 0.83ns | 3.26ns | 3.64ns |
| 16k | 2.14ns | 0.84ns | 20.53ns | 6.09ns | 0.83ns | 3.19ns | 1.39ns |
| 256k | 2.16ns | 0.84ns | 23.93ns | 6.03ns | 0.95ns | 3.20ns | 1.26ns |
| 1M | 2.16ns | 1.00ns | 25.90ns | 6.04ns | 1.94ns | 3.22ns | 1.26ns |


> **Note:** Donburi, GOKe, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](docs/results/random.svg)

| N | Ark | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 2.01ns | 5.33ns | 1.98ns | 6.08ns | 30.41ns | 9.74ns |
| 4 | 1.91ns | 5.18ns | 1.80ns | 6.95ns | 29.80ns | 9.03ns |
| 16 | 1.94ns | 5.20ns | 1.79ns | 8.53ns | 30.14ns | 9.45ns |
| 64 | 1.90ns | 5.31ns | 1.80ns | 8.11ns | 30.90ns | 8.95ns |
| 256 | 1.91ns | 5.44ns | 1.82ns | 8.44ns | 31.79ns | 8.76ns |
| 1k | 1.91ns | 6.78ns | 1.84ns | 9.79ns | 31.72ns | 8.71ns |
| 16k | 2.56ns | 9.95ns | 2.34ns | 11.45ns | 33.43ns | 8.83ns |
| 256k | 7.68ns | 29.23ns | 6.33ns | 32.86ns | 58.14ns | 20.30ns |
| 1M | 7.49ns | 34.17ns | 6.51ns | 37.35ns | 63.74ns | 20.64ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](docs/results/create2comp.svg)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 499.30ns | 451.37ns | 1.53us | 681.06ns | 512.10ns | 742.52ns | 894.14ns |
| 4 | 148.52ns | 134.12ns | 624.42ns | 274.62ns | 152.70ns | 293.12ns | 365.95ns |
| 16 | 67.38ns | 39.41ns | 309.67ns | 161.33ns | 45.65ns | 154.40ns | 168.42ns |
| 64 | 37.31ns | 18.01ns | 246.46ns | 149.00ns | 13.46ns | 100.77ns | 126.93ns |
| 256 | 28.29ns | 10.72ns | 208.57ns | 119.64ns | 5.33ns | 107.27ns | 91.43ns |
| 1k | 28.41ns | 9.01ns | 182.21ns | 111.94ns | 3.81ns | 282.75ns | 78.56ns |
| 16k | 24.81ns | 7.89ns | 197.10ns | 115.48ns | 5.83ns | 286.23ns | 75.01ns |
| 256k | 24.06ns | 7.94ns | 195.04ns | 121.36ns | 7.49ns | 295.87ns | 77.32ns |
| 1M | 24.26ns | 7.93ns | 171.73ns | 174.59ns | 8.37ns | 408.68ns | 77.16ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](docs/results/create2comp_alloc.svg)

| N | Ark | Ark (batch) | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 10.12us | 10.42us | 5.72us | 77.66ns | 18.44us | 4.17us | 2.52us |
| 4 | 3.25us | 3.28us | 1.88us | 54.98ns | 5.04us | 1.42us | 942.81ns |
| 16 | 888.85ns | 848.48ns | 838.64ns | 62.30ns | 1.43us | 452.69ns | 310.65ns |
| 64 | 175.77ns | 159.53ns | 410.88ns | 53.54ns | 599.12ns | 389.87ns | 201.31ns |
| 256 | 99.23ns | 73.73ns | 379.50ns | 42.87ns | 291.13ns | 218.35ns | 152.71ns |
| 1k | 46.90ns | 26.77ns | 310.71ns | 64.39ns | 189.55ns | 182.96ns | 119.54ns |
| 16k | 52.05ns | 35.83ns | 387.57ns | 38.37ns | 258.15ns | 206.43ns | 155.01ns |
| 256k | 50.37ns | 31.21ns | 460.13ns | 67.36ns | 524.97ns | 208.92ns | 125.90ns |
| 1M | 42.73ns | 17.91ns | 371.84ns | 42.35ns | 1.07us | 257.20ns | 111.99ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](docs/results/create10comp.svg)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 407.93ns | 428.00ns | 2.42us | 808.44ns | 587.19ns | 962.20ns | 2.88us |
| 4 | 143.22ns | 139.24ns | 1.21us | 329.91ns | 153.04ns | 474.16ns | 1.64us |
| 16 | 64.72ns | 40.74ns | 976.40ns | 244.05ns | 40.24ns | 287.69ns | 1.36us |
| 64 | 35.85ns | 17.72ns | 783.56ns | 233.15ns | 13.11ns | 254.89ns | 993.45ns |
| 256 | 29.31ns | 10.57ns | 670.85ns | 190.45ns | 6.19ns | 232.35ns | 893.35ns |
| 1k | 28.97ns | 9.27ns | 678.16ns | 169.59ns | 14.59ns | 468.82ns | 863.13ns |
| 16k | 24.72ns | 8.52ns | 685.96ns | 164.93ns | 12.65ns | 474.06ns | 846.91ns |
| 256k | 24.70ns | 7.92ns | 643.31ns | 166.53ns | 16.22ns | 575.96ns | 851.72ns |
| 1M | 24.21ns | 7.95ns | 695.65ns | 240.74ns | 18.07ns | 625.37ns | 855.21ns |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](docs/results/add_remove.svg)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 62.93ns | 119.53ns | 278.18ns | 225.62ns | 323.62ns | 242.79ns | 189.88ns |
| 4 | 66.72ns | 32.31ns | 282.00ns | 228.79ns | 88.14ns | 250.87ns | 192.78ns |
| 16 | 69.61ns | 10.29ns | 283.94ns | 250.55ns | 29.68ns | 259.42ns | 187.92ns |
| 64 | 67.77ns | 5.33ns | 281.60ns | 247.90ns | 14.72ns | 261.36ns | 185.40ns |
| 256 | 68.63ns | 4.31ns | 275.08ns | 252.23ns | 10.49ns | 270.59ns | 183.49ns |
| 1k | 68.45ns | 4.59ns | 276.17ns | 258.97ns | 12.89ns | 674.33ns | 184.00ns |
| 16k | 68.31ns | 4.25ns | 331.14ns | 274.28ns | 28.25ns | 710.26ns | 185.02ns |
| 256k | 71.00ns | 7.08ns | 293.32ns | 403.11ns | 33.28ns | 868.80ns | 178.82ns |
| 1M | 70.60ns | 6.74ns | 302.42ns | 632.80ns | 40.26ns | 1.14us | 182.66ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](docs/results/add_remove_large.svg)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 298.00ns | 323.20ns | 809.34ns | 533.38ns | 492.56ns | 596.16ns | 1.08us |
| 4 | 289.12ns | 86.54ns | 845.38ns | 572.14ns | 132.03ns | 593.82ns | 1.08us |
| 16 | 254.47ns | 38.51ns | 835.70ns | 548.63ns | 46.52ns | 590.17ns | 1.06us |
| 64 | 252.10ns | 24.01ns | 819.98ns | 546.08ns | 25.65ns | 610.71ns | 1.04us |
| 256 | 245.67ns | 18.75ns | 811.72ns | 552.69ns | 32.13ns | 635.90ns | 1.05us |
| 1k | 247.23ns | 17.14ns | 834.93ns | 560.05ns | 68.71ns | 1.13us | 1.05us |
| 16k | 266.86ns | 36.33ns | 873.91ns | 616.28ns | 83.37ns | 1.21us | 1.03us |
| 256k | 297.53ns | 37.07ns | 877.89ns | 777.45ns | 124.05ns | 1.39us | 1.04us |
| 1M | 336.99ns | 69.54ns | 939.01ns | 945.88ns | 179.01ns | 1.71us | 1.04us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](docs/results/delete2comp.svg)

| N | Ark | Ark (batch) | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 162.79ns | 215.37ns | 218.82ns | 750.32ns | 302.02ns | 166.00ns | 241.46ns |
| 4 | 71.18ns | 58.69ns | 93.51ns | 197.71ns | 143.19ns | 72.83ns | 107.98ns |
| 16 | 38.08ns | 18.56ns | 60.09ns | 55.18ns | 125.32ns | 45.39ns | 79.32ns |
| 64 | 32.88ns | 8.51ns | 52.92ns | 21.91ns | 122.34ns | 39.32ns | 74.89ns |
| 256 | 32.54ns | 5.86ns | 51.17ns | 11.83ns | 96.22ns | 40.77ns | 59.94ns |
| 1k | 26.39ns | 5.24ns | 42.67ns | 14.74ns | 105.17ns | 35.13ns | 55.23ns |
| 16k | 23.96ns | 4.71ns | 45.09ns | 14.54ns | 113.09ns | 40.97ns | 53.89ns |
| 256k | 24.17ns | 4.94ns | 49.44ns | 13.50ns | 167.90ns | 60.90ns | 54.02ns |
| 1M | 24.66ns | 5.16ns | 46.15ns | 14.78ns | 284.44ns | 129.32ns | 54.92ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](docs/results/delete10comp.svg)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 203.64ns | 300.12ns | 352.55ns | 404.22ns | 828.92ns | 162.00ns | 576.36ns |
| 4 | 135.27ns | 84.13ns | 189.60ns | 228.88ns | 215.19ns | 69.24ns | 400.05ns |
| 16 | 99.67ns | 27.19ns | 137.12ns | 200.10ns | 63.15ns | 48.46ns | 299.71ns |
| 64 | 105.52ns | 11.62ns | 132.28ns | 173.52ns | 26.06ns | 36.53ns | 243.41ns |
| 256 | 86.50ns | 8.12ns | 117.58ns | 154.25ns | 18.65ns | 43.45ns | 230.98ns |
| 1k | 74.25ns | 7.18ns | 124.52ns | 159.21ns | 32.77ns | 35.12ns | 226.24ns |
| 16k | 72.44ns | 6.59ns | 124.72ns | 179.97ns | 21.22ns | 40.03ns | 224.69ns |
| 256k | 84.65ns | 9.25ns | 147.48ns | 306.40ns | 26.84ns | 78.50ns | 234.34ns |
| 1M | 80.46ns | 9.95ns | 197.26ns | 423.66ns | 46.80ns | 132.66ns | 229.22ns |


### Create world

- Create a new world

| N | Ark | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 21.00us | 2.12us | 903.92us | 205.69us | 2.48us | 17.40us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/ark&yohamta0/donburi-ecs&marioolofo/go-gameengine-ecs&kjkrol/goke&unitoftime/ecs&akmonengine/volt&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/ark,yohamta0/donburi-ecs,marioolofo/go-gameengine-ecs,kjkrol/goke,unitoftime/ecs,akmonengine/volt&type=Date" alt="Star History Chart" width="600"/>
</a>
</p>

## Running the benchmarks

Run the benchmarks using the following command:

```shell
go run . -test.benchtime=0.25s
```

> On PowerShell use this instead:  
> `go run . --% -test.benchtime=0.25s`

The `benchtime` limit is required for some of the benchmarks that have a high
setup cost which is not timed. They would take forever otherwise.
The benchmarks can take up to one hour to complete.

To run a selection of benchmarks, add their names as arguments:

```shell
go run . query2comp query32arch
```

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```

## Contributing

Developers of ECS frameworks are welcome to add their implementation to the benchmarks.
However, there are a few (quality) criteria that need to be fulfilled for inclusion:

- All benchmarks must be implemented, which means that the ECS must have the required features
- The ECS must be working properly and not exhibit serious flaws; it will undergo a basic review by maintainers
- The ECS must be sufficiently documented so that it can be used without reading the code
- There must be at least basic unit tests
- Unit tests must be run in the CI of the repository
- The ECS *must not* be tightly coupled to a particular game engine, particularly graphics stuff
- There must be tagged release versions; only tagged versions will be included here

Developers of included frameworks are encouraged to review the benchmarks,
and to fix (or point to) misuse or potential optimizations.

## Automated results PRs

Pull requests from the `bench-results-update` branch are opened automatically by the `publish` job
in `.github/workflows/benchmarks.yml`, after a benchmark run on `main` — they only refresh
`README.md` and `docs/results/**` (the plotted numbers/images), never benchmark code, and don't need
the code-review criteria above. That branch is force-pushed on every run, so there's at most one such
PR open at a time, reused rather than duplicated. Everything else (`bench/**`, `go.mod`, workflow,
docs) is normal work, reviewed as usual.
