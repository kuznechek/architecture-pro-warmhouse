from typing import Annotated
from fastapi import Depends, FastAPI, HTTPException, Query
from sqlmodel import Session, MetaData, create_engine, select
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from models import Sensor

database_url = 'postgresql://postgres:postgres@postgres/smarthome'

engine = create_engine(database_url)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()


def get_session():
    with Session(engine) as session:
        yield session

SessionDep = Annotated[SessionLocal, Depends(get_session)]
app = FastAPI()

@app.get("/api/device-control-service/")
async def index():
    return {"device control status": "ok"}

@app.post("/api/device-control-service/sensors/")
def create_sensor(sensor: Sensor) -> Sensor:
    session = SessionLocal()
    session.add(sensor)
    session.commit()
    session.refresh(sensor)

    print(f"Sensor was created: {sensor.name}")
    return  sensor

@app.get("/api/device-control-service/sensors/")
def get_sensors(session: SessionDep, offset: int = 0, limit: Annotated[int, Query(le=100)] = 100) -> list[Sensor]:
    sensors = session.exec(select(Sensor).offset(offset).limit(limit)).all()
    return sensors

@app.get("/api/device-control-service/sensors/{sensorId}")
def read_sensor(sensorId: int, session: SessionDep):
    sensor = session.get(Sensor, sensorId)
    if not sensor:
        raise HTTPException(status_code=404, detail="Sensor not found")
    return sensor

@app.delete("/api/device-control-service/sensors/{sensorId}")
def delete_sensor(sensorId: int, session: SessionDep):
    sensor = session.get(Sensor, sensorId)
    if not sensor:
        raise HTTPException(status_code=404, detail="Sensor not found")
    session.delete(sensor)
    session.commit()
    return {"ok": True}
