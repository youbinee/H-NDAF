# H-NDAF

## Introduction
H-NDAF is a hierarchical NWDAF architecture for automated network operation and management in B5G networks. H-NDAF can provide inference results quickly through a new architecture divided into root and leaf NWDAFs.

We implement H-NDAF using free5GC which is a well-known open-source software for 5G CN. 

## Requirements
We recommend the following specifications.

1. Free5GC open-source [Link](https://www.free5gc.org/, "Free5GC")
2. Python 3.5.x  [Link](https://www.python.org/, "Python")
3. TensorFlow 1.10.0 [Link](https://www.tensorflow.org/, "Tensorflow")

H-NDAF works based on the open-source free5GC.
Therefore, the codes should be used after installing the basics of free5GC.

## How to run
1. Register root NWDAF with NRF
```
go run nrf.go
```

2. Run root NWDAF and root learning 
```
go run root_NWDAF.go
python root_learning.py
```

3. Run leaf NWDAF and leaf learning
```
go run leaf_NWDAF.go
python leaf_learning.py
```
