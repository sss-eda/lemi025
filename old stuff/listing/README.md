# Listing Context

This isn't just a "more or less one thing" package. This package contains a single Read Model along with all of the accompanying usecases. This context is specifically aimed at usecases involving "getting" and "listing" usecases, since these will require the same model and database type. (This is opposed to a "searching" usecase, which would most likely require a different kind of database).

There won't be any mutation of state happening in this context. This means that the model won't require an anti-corruption layer.

## Read Model
Below is the design of the read model:

```mermaid
erDiagram
    INSTRUMENT {}
    CONFIG {
        string station-type
        int station-number
    }
    TIME {
        int year
        int month
        int day
        int hour
        int minute
        int second
    }
    COEFFICIENTS1 {
        uint8 mode
        uint8 uin
        uint8 mode1
    }
    COEFFICIENTS2 {
        float32 ax1
        float32 ay1
        float32 az1
        float32 beta
        float32 gamma
        float32 xi
        float32 exy
        float32 eyz
        float32 exz
        float32 k1x
        float32 k1y
        float32 k1z
        float32 k2x
        float32 k2y
        float32 k2z
        float32 ktf
        float32 kte
        float32 ktf0
        float32 kte0
        float32 kvbat
    }
    GPSDATA {
        bytes latitude
        bytes longitude
        bytes altitude
    }
    INSTRUMENT ||--|| CONFIG : contains
    INSTRUMENT ||--|| TIME : contains
    INSTRUMENT ||--|| COEFFICIENTS1 : contains
    INSTRUMENT ||--|| COEFFICIENTS2 : contains
    INSTRUMENT ||--|| GPSDATA : contains
```

## Use Cases
* GetInstrumentByID
* ListInstruments (GetAllInstruments?)