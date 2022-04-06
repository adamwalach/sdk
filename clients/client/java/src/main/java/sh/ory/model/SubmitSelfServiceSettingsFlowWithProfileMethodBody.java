/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.157
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * nolint:deadcode,unused
 */
@ApiModel(description = "nolint:deadcode,unused")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-04-06T19:09:37.881069760Z[Etc/UTC]")
public class SubmitSelfServiceSettingsFlowWithProfileMethodBody {
  public static final String SERIALIZED_NAME_CSRF_TOKEN = "csrf_token";
  @SerializedName(SERIALIZED_NAME_CSRF_TOKEN)
  private String csrfToken;

  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_TRAITS = "traits";
  @SerializedName(SERIALIZED_NAME_TRAITS)
  private Object traits;

  public SubmitSelfServiceSettingsFlowWithProfileMethodBody() { 
  }

  public SubmitSelfServiceSettingsFlowWithProfileMethodBody csrfToken(String csrfToken) {
    
    this.csrfToken = csrfToken;
    return this;
  }

   /**
   * The Anti-CSRF Token  This token is only required when performing browser flows.
   * @return csrfToken
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The Anti-CSRF Token  This token is only required when performing browser flows.")

  public String getCsrfToken() {
    return csrfToken;
  }


  public void setCsrfToken(String csrfToken) {
    this.csrfToken = csrfToken;
  }


  public SubmitSelfServiceSettingsFlowWithProfileMethodBody method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Method  Should be set to profile when trying to update a profile.
   * @return method
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Method  Should be set to profile when trying to update a profile.")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public SubmitSelfServiceSettingsFlowWithProfileMethodBody traits(Object traits) {
    
    this.traits = traits;
    return this;
  }

   /**
   * Traits contains all of the identity&#39;s traits.
   * @return traits
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Traits contains all of the identity's traits.")

  public Object getTraits() {
    return traits;
  }


  public void setTraits(Object traits) {
    this.traits = traits;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SubmitSelfServiceSettingsFlowWithProfileMethodBody submitSelfServiceSettingsFlowWithProfileMethodBody = (SubmitSelfServiceSettingsFlowWithProfileMethodBody) o;
    return Objects.equals(this.csrfToken, submitSelfServiceSettingsFlowWithProfileMethodBody.csrfToken) &&
        Objects.equals(this.method, submitSelfServiceSettingsFlowWithProfileMethodBody.method) &&
        Objects.equals(this.traits, submitSelfServiceSettingsFlowWithProfileMethodBody.traits);
  }

  @Override
  public int hashCode() {
    return Objects.hash(csrfToken, method, traits);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SubmitSelfServiceSettingsFlowWithProfileMethodBody {\n");
    sb.append("    csrfToken: ").append(toIndentedString(csrfToken)).append("\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    traits: ").append(toIndentedString(traits)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

