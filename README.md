# geojson

> Geojson is a package for creating geojson using golang language


## Go get
> go get github.com/Mziia/geojson/...


## Quick start

```
collection := features.NewFeatureCollection() 

point := coordinates.NewPoint([]float64{125.6, 10.1}) 

feature := features.NewFeature()  
                    
feature.SetProperty("id_num", 1)
feature.SetID("0001")
feature.SetGeometry(point)
	
collection.AddFeature(feature)
```
