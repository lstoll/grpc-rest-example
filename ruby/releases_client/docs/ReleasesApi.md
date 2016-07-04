# ReleasesClient::ReleasesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_release**](ReleasesApi.md#add_release) | **POST** /v1/release | 
[**get_channel_releases**](ReleasesApi.md#get_channel_releases) | **GET** /v1/channel_releases/{ChannelID} | 
[**get_release**](ReleasesApi.md#get_release) | **GET** /v1/release/{ReleaseID} | 


# **add_release**
> ReleaserAddReleaseResponse add_release(body)



### Example
```ruby
# load the gem
require 'releases_client'

api_instance = ReleasesClient::ReleasesApi.new

body = ReleasesClient::ReleaserRelease.new # ReleaserRelease | 


begin
  result = api_instance.add_release(body)
  p result
rescue ReleasesClient::ApiError => e
  puts "Exception when calling ReleasesApi->add_release: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ReleaserRelease**](ReleaserRelease.md)|  | 

### Return type

[**ReleaserAddReleaseResponse**](ReleaserAddReleaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_channel_releases**
> ReleaserChannelReleases get_channel_releases(channel_id)



### Example
```ruby
# load the gem
require 'releases_client'

api_instance = ReleasesClient::ReleasesApi.new

channel_id = "channel_id_example" # String | 


begin
  result = api_instance.get_channel_releases(channel_id)
  p result
rescue ReleasesClient::ApiError => e
  puts "Exception when calling ReleasesApi->get_channel_releases: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel_id** | **String**|  | 

### Return type

[**ReleaserChannelReleases**](ReleaserChannelReleases.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_release**
> ReleaserRelease get_release(release_id)



### Example
```ruby
# load the gem
require 'releases_client'

api_instance = ReleasesClient::ReleasesApi.new

release_id = "release_id_example" # String | 


begin
  result = api_instance.get_release(release_id)
  p result
rescue ReleasesClient::ApiError => e
  puts "Exception when calling ReleasesApi->get_release: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **release_id** | **String**|  | 

### Return type

[**ReleaserRelease**](ReleaserRelease.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



