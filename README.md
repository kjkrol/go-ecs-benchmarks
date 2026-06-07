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
| [GOKe](https://github.com/kjkrol/goke) | ✅ | ❌ | ❌ | ❌ | ✅ | ✅ |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |
| [Volt](https://github.com/akmonengine/volt) | ✅ | ❌ | ❌ | ✅ | ❌ | ❌ |

[1] ECS lifecycle events, allowing to react to entity creation, component addition, ...  
[2] Faster batch operations for entity creation etc.

## Benchmarks

Last run: Mon, 01 Jun 2026 11:47:54 UTC  
CPU: Intel(R) Xeon(R) Platinum 8370C CPU @ 2.80GHz


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

![query2comp](https://github.com/user-attachments/assets/ff4cb720-2660-4432-8bab-eeba93706c25)

| N | Ark | Ark (tables) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 64.57ns | 70.10ns | 63.39ns | 48.17ns | 7.34ns | 18.27ns | 79.12ns |
| 4 | 17.87ns | 18.49ns | 29.63ns | 15.55ns | 2.82ns | 6.88ns | 20.57ns |
| 16 | 5.74ns | 5.35ns | 22.17ns | 7.40ns | 1.47ns | 4.58ns | 6.10ns |
| 64 | 2.67ns | 2.00ns | 19.29ns | 5.72ns | 0.84ns | 3.46ns | 2.48ns |
| 256 | 1.95ns | 1.05ns | 19.92ns | 5.19ns | 0.70ns | 3.22ns | 1.60ns |
| 1k | 1.82ns | 0.90ns | 19.92ns | 5.06ns | 0.68ns | 3.27ns | 1.37ns |
| 16k | 1.78ns | 0.84ns | 21.12ns | 5.05ns | 0.81ns | 3.26ns | 1.28ns |
| 256k | 1.77ns | 0.83ns | 22.89ns | 5.09ns | 0.80ns | 3.20ns | 1.26ns |
| 1M | 1.84ns | 0.96ns | 30.63ns | 5.05ns | 1.07ns | 3.20ns | 1.35ns |


> **Note:** Donburi, GOKe, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/ff6a38fb-72cd-4072-9fb9-186181b2652c)

| N | Ark | Ark (tables) | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 65.25ns | 71.57ns | 62.71ns | 7.40ns | 45.74ns | 17.19ns | 83.12ns |
| 4 | 32.42ns | 39.13ns | 32.91ns | 5.11ns | 19.04ns | 14.89ns | 67.25ns |
| 16 | 23.18ns | 31.76ns | 26.46ns | 4.51ns | 12.45ns | 24.54ns | 62.06ns |
| 64 | 13.08ns | 16.25ns | 23.17ns | 3.43ns | 8.51ns | 14.55ns | 32.74ns |
| 256 | 4.92ns | 4.80ns | 21.20ns | 2.68ns | 5.76ns | 6.10ns | 9.37ns |
| 1k | 2.67ns | 1.87ns | 22.71ns | 1.26ns | 5.21ns | 3.99ns | 3.47ns |
| 16k | 1.95ns | 0.99ns | 29.90ns | 0.90ns | 5.08ns | 3.29ns | 1.54ns |
| 256k | 1.78ns | 0.83ns | 67.57ns | 0.82ns | 5.05ns | 3.18ns | 1.28ns |
| 1M | 1.84ns | 1.08ns | 95.13ns | 1.23ns | 5.06ns | 3.20ns | 1.37ns |


> **Note:** Donburi, GOKe, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/83e0bb70-24d7-4d98-a2eb-64f010b1315f)

| N | Ark | Ark (tables) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 62.22ns | 68.11ns | 63.51ns | 55.53ns | 21.48ns | 17.18ns | 100.59ns |
| 4 | 17.44ns | 17.50ns | 29.59ns | 25.96ns | 18.81ns | 9.72ns | 47.94ns |
| 16 | 5.69ns | 4.96ns | 21.19ns | 19.01ns | 25.46ns | 4.73ns | 36.19ns |
| 64 | 2.73ns | 1.93ns | 19.45ns | 16.96ns | 29.94ns | 3.79ns | 37.37ns |
| 256 | 2.01ns | 1.07ns | 20.15ns | 8.13ns | 12.01ns | 3.38ns | 10.51ns |
| 1k | 1.91ns | 0.89ns | 20.08ns | 5.80ns | 9.16ns | 3.17ns | 3.57ns |
| 16k | 1.93ns | 0.82ns | 20.82ns | 5.11ns | 3.91ns | 3.18ns | 1.45ns |
| 256k | 1.96ns | 0.80ns | 23.61ns | 5.06ns | 6.24ns | 3.17ns | 1.28ns |
| 1M | 2.03ns | 1.17ns | 27.19ns | 5.10ns | 6.93ns | 3.20ns | 1.41ns |


> **Note:** Donburi, GOKe, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/ffac498e-a9bd-400c-a121-247b75af86f3)

| N | Ark | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 3.03ns | 9.42ns | 16.55ns | 9.32ns | 39.12ns | 13.73ns |
| 4 | 2.87ns | 8.99ns | 9.12ns | 9.25ns | 38.88ns | 13.49ns |
| 16 | 2.96ns | 8.57ns | 6.22ns | 14.28ns | 39.58ns | 12.89ns |
| 64 | 3.05ns | 8.65ns | 5.73ns | 14.42ns | 42.86ns | 12.99ns |
| 256 | 3.18ns | 9.83ns | 6.22ns | 16.21ns | 43.74ns | 13.26ns |
| 1k | 2.99ns | 15.08ns | 6.03ns | 19.51ns | 43.06ns | 13.52ns |
| 16k | 6.20ns | 45.13ns | 12.88ns | 36.85ns | 68.22ns | 17.86ns |
| 256k | 9.15ns | 165.38ns | 19.22ns | 118.57ns | 164.13ns | 25.43ns |
| 1M | 48.20ns | 277.23ns | 85.79ns | 180.91ns | 253.13ns | 118.56ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/5854ec2b-65b5-4df9-acb0-df92aa4d6aeb)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 235.14ns | 263.06ns | 1.25us | 442.41ns | 254.34ns | 488.68ns | 726.32ns |
| 4 | 86.04ns | 73.56ns | 503.18ns | 175.09ns | 100.07ns | 207.12ns | 268.66ns |
| 16 | 46.61ns | 25.41ns | 328.37ns | 142.33ns | 67.66ns | 132.84ns | 148.35ns |
| 64 | 35.31ns | 12.73ns | 274.88ns | 133.48ns | 60.88ns | 107.74ns | 135.57ns |
| 256 | 30.88ns | 10.09ns | 205.76ns | 111.91ns | 54.11ns | 99.04ns | 92.58ns |
| 1k | 30.39ns | 10.29ns | 191.49ns | 110.87ns | 49.71ns | 360.79ns | 87.01ns |
| 16k | 23.06ns | 8.21ns | 207.83ns | 117.70ns | 51.25ns | 352.05ns | 84.02ns |
| 256k | 22.70ns | 8.25ns | 216.70ns | 132.35ns | 50.38ns | 412.51ns | 83.55ns |
| 1M | 22.83ns | 8.43ns | 183.99ns | 226.05ns | 50.75ns | 465.49ns | 83.62ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/e0c24634-c58e-4fa4-a9b0-ef6878bd1278)

| N | Ark | Ark (batch) | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 8.77us | 8.83us | 5.55us | 694.23ns | 21.77us | 3.53us | 1.52us |
| 4 | 2.26us | 2.32us | 1.64us | 225.30ns | 4.95us | 1.21us | 655.55ns |
| 16 | 601.53ns | 564.16ns | 656.56ns | 110.01ns | 1.68us | 409.25ns | 275.95ns |
| 64 | 176.66ns | 154.87ns | 435.00ns | 72.27ns | 598.60ns | 227.93ns | 169.58ns |
| 256 | 71.05ns | 48.73ns | 329.15ns | 62.45ns | 274.29ns | 181.97ns | 151.35ns |
| 1k | 42.87ns | 22.29ns | 311.29ns | 60.22ns | 240.93ns | 166.74ns | 125.64ns |
| 16k | 51.97ns | 43.27ns | 432.03ns | 58.67ns | 257.71ns | 208.14ns | 152.06ns |
| 256k | 63.64ns | 33.21ns | 490.25ns | 61.21ns | 590.99ns | 213.69ns | 146.03ns |
| 1M | 41.78ns | 21.55ns | 438.47ns | 55.87ns | 1.13us | 271.69ns | 122.60ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/29f5b855-0a04-404c-994c-a82e015bf35c)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 271.74ns | 231.29ns | 2.05us | 552.75ns | 288.18ns | 755.48ns | 2.69us |
| 4 | 94.90ns | 64.02ns | 1.29us | 244.37ns | 134.83ns | 397.72ns | 1.52us |
| 16 | 45.79ns | 22.55ns | 1.07us | 213.33ns | 100.37ns | 276.22ns | 1.22us |
| 64 | 35.34ns | 12.76ns | 846.41ns | 221.80ns | 89.94ns | 261.02ns | 967.33ns |
| 256 | 30.04ns | 10.19ns | 732.01ns | 162.32ns | 78.07ns | 209.06ns | 911.42ns |
| 1k | 30.64ns | 9.21ns | 733.47ns | 165.32ns | 105.08ns | 477.30ns | 906.52ns |
| 16k | 23.03ns | 7.91ns | 744.09ns | 167.40ns | 90.14ns | 485.23ns | 919.34ns |
| 256k | 25.03ns | 8.02ns | 730.01ns | 186.51ns | 77.34ns | 597.62ns | 879.88ns |
| 1M | 23.57ns | 7.92ns | 755.32ns | 254.82ns | 78.66ns | 712.05ns | 884.82ns |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/b0687768-23a2-48d2-8ede-6a85b4607c35)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 94.26ns | 175.26ns | 424.22ns | 551.07ns | 100.93ns | 346.14ns | 269.01ns |
| 4 | 103.06ns | 46.05ns | 434.91ns | 517.98ns | 117.61ns | 368.82ns | 288.90ns |
| 16 | 116.21ns | 22.96ns | 461.93ns | 576.86ns | 115.58ns | 379.05ns | 275.31ns |
| 64 | 119.30ns | 13.88ns | 478.17ns | 586.26ns | 130.25ns | 390.33ns | 281.29ns |
| 256 | 113.32ns | 9.47ns | 500.93ns | 634.93ns | 133.44ns | 442.01ns | 346.19ns |
| 1k | 121.67ns | 9.69ns | 484.85ns | 606.05ns | 111.59ns | 776.23ns | 273.92ns |
| 16k | 107.45ns | 10.08ns | 510.57ns | 618.32ns | 124.08ns | 847.29ns | 272.76ns |
| 256k | 120.70ns | 12.16ns | 519.11ns | 919.32ns | 120.93ns | 1.31us | 282.73ns |
| 1M | 115.79ns | 12.08ns | 497.45ns | 1.08us | 126.40ns | 1.46us | 269.10ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/1a166f99-61e1-4414-9f76-488ddf8202e3)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 344.97ns | 459.63ns | 1.10us | 969.83ns | 271.90ns | 803.15ns | 1.50us |
| 4 | 401.62ns | 122.81ns | 1.14us | 988.26ns | 372.31ns | 823.58ns | 1.41us |
| 16 | 405.94ns | 47.65ns | 1.09us | 1.05us | 356.68ns | 812.59ns | 1.41us |
| 64 | 412.96ns | 28.32ns | 1.09us | 1.06us | 356.92ns | 879.76ns | 1.41us |
| 256 | 409.67ns | 22.54ns | 1.11us | 1.07us | 392.21ns | 994.47ns | 1.61us |
| 1k | 431.39ns | 20.96ns | 1.24us | 1.01us | 425.42ns | 1.51us | 1.46us |
| 16k | 423.38ns | 26.64ns | 1.27us | 1.24us | 500.33ns | 1.78us | 1.60us |
| 256k | 593.08ns | 45.17ns | 1.52us | 1.63us | 445.56ns | 2.16us | 1.65us |
| 1M | 495.37ns | 44.80ns | 1.58us | 1.85us | 523.10ns | 2.48us | 1.87us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/d23a5a87-2e4e-46fc-ae14-c79143dcab1b)

| N | Ark | Ark (batch) | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 140.62ns | 243.47ns | 221.33ns | 166.36ns | 340.23ns | 164.53ns | 290.58ns |
| 4 | 72.64ns | 66.18ns | 94.26ns | 90.07ns | 172.72ns | 73.49ns | 121.19ns |
| 16 | 46.11ns | 21.11ns | 72.01ns | 65.92ns | 157.56ns | 54.96ns | 101.00ns |
| 64 | 37.75ns | 8.73ns | 57.01ns | 59.94ns | 122.81ns | 46.59ns | 76.73ns |
| 256 | 28.45ns | 7.30ns | 45.81ns | 43.72ns | 110.68ns | 40.08ns | 53.74ns |
| 1k | 28.61ns | 5.66ns | 38.58ns | 29.63ns | 93.75ns | 32.88ns | 55.87ns |
| 16k | 22.62ns | 4.04ns | 42.36ns | 30.37ns | 101.84ns | 37.76ns | 54.01ns |
| 256k | 23.27ns | 4.98ns | 42.23ns | 31.82ns | 241.90ns | 80.78ns | 51.83ns |
| 1M | 25.69ns | 6.24ns | 55.25ns | 33.51ns | 323.32ns | 129.66ns | 56.58ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/c1df1169-38c3-430d-9c15-d1df391d4323)

| N | Ark | Ark (batch) | Donburi | ggecs | GOKe | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 197.78ns | 303.94ns | 402.98ns | 483.48ns | 232.17ns | 154.54ns | 540.20ns |
| 4 | 136.65ns | 80.35ns | 200.80ns | 297.14ns | 162.98ns | 72.85ns | 455.88ns |
| 16 | 147.44ns | 28.79ns | 171.13ns | 260.05ns | 147.26ns | 48.67ns | 322.51ns |
| 64 | 100.46ns | 14.86ns | 134.78ns | 169.22ns | 93.72ns | 49.54ns | 234.96ns |
| 256 | 75.88ns | 8.60ns | 107.21ns | 142.86ns | 89.24ns | 42.57ns | 212.31ns |
| 1k | 69.26ns | 7.61ns | 104.43ns | 144.76ns | 85.37ns | 34.32ns | 289.04ns |
| 16k | 66.49ns | 9.91ns | 107.97ns | 166.74ns | 77.57ns | 33.49ns | 206.58ns |
| 256k | 75.69ns | 13.74ns | 150.44ns | 342.90ns | 88.27ns | 78.40ns | 207.33ns |
| 1M | 71.36ns | 16.68ns | 238.52ns | 429.94ns | 92.04ns | 122.24ns | 211.13ns |


### Create world

- Create a new world

| N | Ark | Donburi | GOKe | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 30.52us | 3.34us | 500.12us | 307.36us | 3.24us | 28.65us |


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
