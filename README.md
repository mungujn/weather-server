# Weather app backend

[![Build Status](https://travis-ci.com/mungujn/weather-server.svg?branch=master)](https://travis-ci.com/mungujn/weather-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/mungujn/weather-server)](https://goreportcard.com/report/github.com/mungujn/weather-server)

Backend for a weather app. App consists of 3 services; an api, a database service and the core weather service.

## API Service

This service exposes an API endpoint that client apps access to retrieve information on the weather in a given location, at a particular time.

## Weather service

This service retrieves weather information from a provider. The server for this service is an RPC server, built using gRPC. It exposes one RPC service, 'GetWeather' for retrieving the weather. The GetWeather operation first uses an RPC client (that connects to the database RPC service) to read weather data cached in the database.
If the weather request is for a location not in the database, or if the weather data retrieved from the database is too old, the service fetches fresh data from a weather provider (still not yet determined), immediately returns this data to the API service and then uses a separate thread (goroutine) to cache the newly retrieved weather data by making a create RPC call to the database service

## Database service

This service is also a remote procedure call server. It exposes 4 methods, for the standard database CRUD operations.

Communication between services is encrypted and secured with SSL/TLS.