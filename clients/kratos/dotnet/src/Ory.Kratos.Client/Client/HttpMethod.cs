/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


namespace Ory.Kratos.Client.Client
{
    /// <summary>
    /// Http methods supported by swagger
    /// </summary>
    public enum HttpMethod
    {
        /// <summary>HTTP GET request.</summary>
        Get,
        /// <summary>HTTP POST request.</summary>
        Post,
        /// <summary>HTTP PUT request.</summary>
        Put,
        /// <summary>HTTP DELETE request.</summary>
        Delete,
        /// <summary>HTTP HEAD request.</summary>
        Head,
        /// <summary>HTTP OPTIONS request.</summary>
        Options,
        /// <summary>HTTP PATCH request.</summary>
        Patch
    }
}
