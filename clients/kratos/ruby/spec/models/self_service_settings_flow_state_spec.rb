=begin
#Ory Kratos API

#Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 

The version of the OpenAPI document: v0.7.0-alpha.1
Contact: hi@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.1.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryHydraClient::SelfServiceSettingsFlowState
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryHydraClient::SelfServiceSettingsFlowState do
  let(:instance) { OryHydraClient::SelfServiceSettingsFlowState.new }

  describe 'test an instance of SelfServiceSettingsFlowState' do
    it 'should create an instance of SelfServiceSettingsFlowState' do
      expect(instance).to be_instance_of(OryHydraClient::SelfServiceSettingsFlowState)
    end
  end
end
