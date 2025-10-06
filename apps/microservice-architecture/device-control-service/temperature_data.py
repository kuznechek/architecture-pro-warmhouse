import requests

# api-endpoint
URL = "http://temperature-api:8081/temperature"

# def GetTemperatureFromExternalService(sensor_id):
#     result = requests.get(url = f'{URL}/{sensor_id}')
#     data = result.json()
#     print(data)

def GetTemperature():
    return GetTemperatureFromExternalService()

def GetTemperatureFromExternalService():
    result = requests.get(url = f'{URL}/')
    data = result.json()
    return data['value']