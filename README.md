# Workshop materials

examples of the implementation of generative art in GO

## What is this?

Based on the original image and using a random selection of coordinates, taking the color at this particular point, 
the program, applying a given finite number of processing steps,
thus transforms and translates the designated color information into the destination space and draws pre-defined geometric form, 
generating a pop art style sketch.

## Examples

![source.jpg](assets%2Fsource.jpg)
![result.png](assets%2Fresult.png)

## Build & Run

```bash
go build -o popart ./cmd
./popart source destination  
```