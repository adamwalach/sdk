/*
 * Ory Kratos API
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.0-alpha.1
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.kratos.model;

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
import java.util.ArrayList;
import java.util.List;
import sh.ory.kratos.model.PluginInterfaceType;

/**
 * PluginConfigInterface The interface between Docker and the plugin
 */
@ApiModel(description = "PluginConfigInterface The interface between Docker and the plugin")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2021-07-13T18:51:24.300092764Z[Etc/UTC]")
public class PluginConfigInterface {
  public static final String SERIALIZED_NAME_SOCKET = "Socket";
  @SerializedName(SERIALIZED_NAME_SOCKET)
  private String socket;

  public static final String SERIALIZED_NAME_TYPES = "Types";
  @SerializedName(SERIALIZED_NAME_TYPES)
  private List<PluginInterfaceType> types = new ArrayList<>();


  public PluginConfigInterface socket(String socket) {
    
    this.socket = socket;
    return this;
  }

   /**
   * socket
   * @return socket
  **/
  @ApiModelProperty(required = true, value = "socket")

  public String getSocket() {
    return socket;
  }


  public void setSocket(String socket) {
    this.socket = socket;
  }


  public PluginConfigInterface types(List<PluginInterfaceType> types) {
    
    this.types = types;
    return this;
  }

  public PluginConfigInterface addTypesItem(PluginInterfaceType typesItem) {
    this.types.add(typesItem);
    return this;
  }

   /**
   * types
   * @return types
  **/
  @ApiModelProperty(required = true, value = "types")

  public List<PluginInterfaceType> getTypes() {
    return types;
  }


  public void setTypes(List<PluginInterfaceType> types) {
    this.types = types;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PluginConfigInterface pluginConfigInterface = (PluginConfigInterface) o;
    return Objects.equals(this.socket, pluginConfigInterface.socket) &&
        Objects.equals(this.types, pluginConfigInterface.types);
  }

  @Override
  public int hashCode() {
    return Objects.hash(socket, types);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PluginConfigInterface {\n");
    sb.append("    socket: ").append(toIndentedString(socket)).append("\n");
    sb.append("    types: ").append(toIndentedString(types)).append("\n");
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

