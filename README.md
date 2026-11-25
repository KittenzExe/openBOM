![](/assets/openbom.png)

# openBOM

openBOM simplifies the API from The Bureau of Meteorology website, allowing developers in Australia to access weather data more easily.

## API

`GET openbom.exe.fyi/v1/location/search?q={location}`<br>
Returns a list of location candidates matching the query.

`GET openbom.exe.fyi/v1/weather/hourly/query?q={location}`<br>
Returns the hourly weather forecast for the queried location.

`GET openbom.exe.fyi/v1/weather/daily/query?q={location}`<br>
Returns the daily weather forecast for the queried location.

`GET openbom.exe.fyi/v1/weather/astro/query?q={location}`<br>
Returns the astronomical weather data for the queried location.

`GET openbom.exe.fyi/v1/weather/hourly/direct?x={x}&y={y}&t={timezone}`<br>
Returns the hourly weather forecast for the specified grid cell.

`GET openbom.exe.fyi/v1/weather/daily/direct?x={x}&y={y}&t={timezone}`<br>
Returns the daily weather forecast for the specified grid cell.

`GET openbom.exe.fyi/v1/weather/astro/direct?lat={lat}&lon={lon}`<br>
Returns the astronomical weather data for the specified coordinates.

## API Usage Examples

>[!NOTE]
> Any time-related weather responses are in UTC-0, you need to adjust to your local timezone.

As an example, say if you want to search for a weather station:

```
GET openbom.exe.fyi/v1/location/search?q=Perth

{
    "candidates": [
        {
            "id": "bwa_pt053",
            "coordinate": {
                "longitude": 115.8587,
                "latitude": -31.95184
            },
            "gridcells": {
                "forecast": {
                    "x": 65,
                    "y": 261
                }
            },
            "name": "Perth",
            "state": "WA",
            "postcode": {
                "name": "6000",
                "description": null
            },
            "timezone": "Australia/Perth",
            "type": "place",
            "feature_code": "place",
            "is_alpine": false
        }
    ...
```

Hourly data? That can be queried as well:

```
GET openbom.exe.fyi/v1/weather/hourly/query?q=Perth

{
    "meta": {
        "issue_time_utc": "2026-11-25T19:50:01Z",
        "issue_time_next_utc": "2026-11-26T06:00:00Z",
        "local_timezone": "Australia/Perth"
    },
    "fcst": [
        {
            "date_utc": "2026-11-25T16:00:00Z",
            "1hourly": [
                {
                    "time_utc": "2026-11-25T19:00:00Z",
                    "atm": {
                        "surf_air": {
                            "temp_apparent_cel": 11.2,
                            "temp_dew_pt_cel": 9.900001,
                            "temp_cel": 12.5,
                            "hum_relative_percent": 84.2,
                            "wind": {
                                "hazardous_onset_code": 3,
                                "dirn_10m_deg_t": 155.8,
                                "speed_10m_avg_mps": 1.944,
                                "speed_10m_avg_kts": 4.0,
                                "gust_speed_10m_max_mps": 3.601,
                                "gust_speed_10m_max_kts": 7.0,
                                "mixing_height_m": 11.0
                            },
                            "radiation": {
                                "uv_clear_sky_code": 0.0
                            }
                        }
                    }
                },
            ...
```

Daily data can also be fetched as easily as hourly data:

```
GET openbom.exe.fyi/v1/weather/daily/query?q=Perth

{
    "meta": {
        "issue_time_utc": "2026-11-25T18:00:22Z",
        "issue_time_next_utc": "2026-11-26T07:00:00Z",
        "local_timezone": "Australia/Perth"
    },
    "fcst": {
        "daily": [
            {
                "date_utc": "2026-11-25T16:00:00Z",
                "atm": {
                    "surf_air": {
                        "temp_max_cel": 25.2,
                        "temp_min_cel": 11.900001,
                        "precip": {
                            "exceeding_10percentchance_total_mm": 0.0,
                            "exceeding_25percentchance_total_mm": 0.0,
                            "exceeding_50percentchance_total_mm": 0.0,
                            "exceeding_75percentchance_total_mm": 0.0,
                            "any_probability_percent": 1.0,
                            "any_restofday_probability_percent": 1.0,
                            "10mm_probability_percent": 0.0,
                            "25mm_probability_percent": 0.0
                        },
                        "weather": {
                            "icon_code": 1
                        },
                        "radiation": {
                            "uv_clear_sky_max_code": 11.86695,
                            "uv_period_start": "2026-11-25T23:50:00Z",
                            "uv_period_end": "2026-11-26T08:10:00Z"
                        }
                    }
                },
                "terr": {
                    "surf_land": {
                        "snow": {}
                    }
                },
                "ocn": {
                    "surf_water": {
                        "sea": {}
                    }
                },
                "astro": {
                    "sunrise_utc": null,
                    "sunset_utc": null
                }
            },
        ...
```


Astronomical data is also available:

```
GET openbom.exe.fyi/v1/weather/astro/query?q=Perth

{
    "meta": {
        "local_timezone": "Australia/Perth"
    },
    "fcst": {
        "daily": [
            {
                "date_utc": "2026-11-25T16:00:00Z",
                "astro": {
                    "sunrise_utc": "2026-11-25T21:04:17Z",
                    "sunset_utc": "2026-11-26T11:03:38Z"
                }
            },
        ...
```

Need a direct spot? You can query by grid cells as well (`lat` and `lon` without a timezone for astronomical data):

```
GET openbom.exe.fyi/v1/weather/hourly/direct?x=65&y=261&t=Australia/Perth

{
    "meta": {
        "issue_time_utc": "2026-11-25T19:50:01Z",
        "issue_time_next_utc": "2026-11-26T06:00:00Z",
        "local_timezone": "Australia/Perth"
    },
    "fcst": [
        {
            "date_utc": "2026-11-25T16:00:00Z",
            "1hourly": [
                {
                    "time_utc": "2026-11-25T19:00:00Z",
                    "atm": {
                        "surf_air": {
                            "temp_apparent_cel": 11.2,
                            "temp_dew_pt_cel": 9.900001,
                            "temp_cel": 12.5,
                            "hum_relative_percent": 84.2,
                            "wind": {
                                "hazardous_onset_code": 3,
                                "dirn_10m_deg_t": 155.8,
                                "speed_10m_avg_mps": 1.944,
                                "speed_10m_avg_kts": 4.0,
                                "gust_speed_10m_max_mps": 3.601,
                                "gust_speed_10m_max_kts": 7.0,
                                "mixing_height_m": 11.0
                            },
                            "radiation": {
                                "uv_clear_sky_code": 0.0
                            }
                        }
                    }
                },
            ...
```

## Run (Local)

To run the project locally (requires Go 1.27.3):

```bash
go build -o openbom .
./openbom
```

## Run (Docker)

To run the project using Docker:

```bash
docker build -t openbom .
docker run -p 3838:3838 openbom
```

## Attribution

Weather data © Commonwealth of Australia (Bureau of Meteorology).

Where applicable, data is licensed by the Commonwealth of Australia under the Creative Commons Attribution 4.0 International License.

(Full disclaimer can be found in the [disclaimer.md](disclaimer.md) file or at [openbom.exe.fyi](https://openbom.exe.fyi).)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.