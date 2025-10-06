from sqlmodel import Field, SQLModel
from datetime import datetime

class Sensor(SQLModel, table=True):
    __tablename__ = 'sensors'
    
    id: int | None = Field(default=None, primary_key=True)
    name: str = Field(index=True)
    type: str = Field(index=True)
    location: str = Field(index=True)
    value: float | None = Field(default=0.0)
    unit: str | None = Field(default="°C")
    status: str = Field(default='inactive')
    created_at: str = Field(default=datetime.now().strftime("%Y-%m-%dT%H:%M:%S%z"))
    last_updated: str = Field(default=datetime.now().strftime("%Y-%m-%dT%H:%M:%S%z"))
