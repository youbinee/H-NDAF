# H-NDAF

## Introduction
H-NDAF is a hierarchical NWDAF architecture for automated network operation and management in B5G networks. H-NDAF can provide inference results quickly through a new architecture divided into root and leaf NWDAFs.

We implement H-NDAF using free5GC which is a well-known open-source software for 5G CN. 

## Requirements
We recommend the following specifications.

1. Free5GC open-source [Link](https://www.free5gc.org/)
2. Python 3.5.x  [Link](https://www.python.org/)
3. TensorFlow 1.10.0 [Link](https://www.tensorflow.org/)

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

## Use Cases
We describe a use case of how PCF in H-NDAF exploits the analytics on the UE downlink throughput prediction (i.e., the analytics on how the achievable data rate of UE changes) to optimize the policy for UE data flows.

H-NDAF can proceed with various additional use cases for analytics in PCF.

**1. Checking Voice Call Bandwidth**

When a UE initiates a voice call (e.g., VoLTE, VoNR, VoWiFi), PCF can request UE bandwidth prediction information from the leaf NWDAF. The root NWDAF generates the UE bandwidth prediction model using information such as UE bandwidth, UE mobility, and serving cell information (e.g., signal strength, signal-to-noise ratio). PCF allocates additional bandwidth required for the UE using the analytics information. This ensures meeting the bandwidth requirements while maintaining the voice call quality for the UE.

**2. Selecting RAN (LTE/NR/WiFi)**

PCF can request analytics on UE network status. The root NWDAF generates a UE network status analytics model using information such as UE location, RAN status, and signal strength. PCF utilizes the analytics information to choose a specific RAN (e.g., LTE, NR, WiFi) and sets policies accordingly.

**3. Setting Traffic Distribution Ratio**

PCF requests a model to analyze the current network traffic situation of LTE and NR. The root NWDAF creates a model based on the traffic data of LTE/NR and communicates it to the leaf NWDAF. PCF uses that analytics to determine how congested RAN is currently and how bandwidth is distributed. This allows PCF to decide how to allocate voice call traffic. For example, if LTE is congested, it can be adjusted to allocate more traffic to NR.
