/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// AWSFlavourBuilder contains the data and logic needed to build 'AWS_flavour' objects.
//
// Volume specification for different classes of nodes inside a flavour.
type AWSFlavourBuilder struct {
	infraVolume  *AWSVolumeBuilder
	masterVolume *AWSVolumeBuilder
	workerVolume *AWSVolumeBuilder
}

// NewAWSFlavour creates a new builder of 'AWS_flavour' objects.
func NewAWSFlavour() *AWSFlavourBuilder {
	return new(AWSFlavourBuilder)
}

// InfraVolume sets the value of the 'infra_volume' attribute
// to the given value.
//
// Holds settings for an AWS storage volume.
func (b *AWSFlavourBuilder) InfraVolume(value *AWSVolumeBuilder) *AWSFlavourBuilder {
	b.infraVolume = value
	return b
}

// MasterVolume sets the value of the 'master_volume' attribute
// to the given value.
//
// Holds settings for an AWS storage volume.
func (b *AWSFlavourBuilder) MasterVolume(value *AWSVolumeBuilder) *AWSFlavourBuilder {
	b.masterVolume = value
	return b
}

// WorkerVolume sets the value of the 'worker_volume' attribute
// to the given value.
//
// Holds settings for an AWS storage volume.
func (b *AWSFlavourBuilder) WorkerVolume(value *AWSVolumeBuilder) *AWSFlavourBuilder {
	b.workerVolume = value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AWSFlavourBuilder) Copy(object *AWSFlavour) *AWSFlavourBuilder {
	if object == nil {
		return b
	}
	if object.infraVolume != nil {
		b.infraVolume = NewAWSVolume().Copy(object.infraVolume)
	} else {
		b.infraVolume = nil
	}
	if object.masterVolume != nil {
		b.masterVolume = NewAWSVolume().Copy(object.masterVolume)
	} else {
		b.masterVolume = nil
	}
	if object.workerVolume != nil {
		b.workerVolume = NewAWSVolume().Copy(object.workerVolume)
	} else {
		b.workerVolume = nil
	}
	return b
}

// Build creates a 'AWS_flavour' object using the configuration stored in the builder.
func (b *AWSFlavourBuilder) Build() (object *AWSFlavour, err error) {
	object = new(AWSFlavour)
	if b.infraVolume != nil {
		object.infraVolume, err = b.infraVolume.Build()
		if err != nil {
			return
		}
	}
	if b.masterVolume != nil {
		object.masterVolume, err = b.masterVolume.Build()
		if err != nil {
			return
		}
	}
	if b.workerVolume != nil {
		object.workerVolume, err = b.workerVolume.Build()
		if err != nil {
			return
		}
	}
	return
}
