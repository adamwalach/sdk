/*
 * Ory Hydra API
 *
 * Documentation for all of Ory Hydra's APIs. 
 *
 * The version of the OpenAPI document: v2.0.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct GetVersion200Response {
    /// The version of Ory Hydra.
    #[serde(rename = "version", skip_serializing_if = "Option::is_none")]
    pub version: Option<String>,
}

impl Default for GetVersion200Response {
    fn default() -> Self {
        Self::new()
    }
}

impl GetVersion200Response {
    pub fn new() -> GetVersion200Response {
        GetVersion200Response {
                version: None,
        }
    }
}


