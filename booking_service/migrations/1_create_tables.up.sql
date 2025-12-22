CREATE TABLE booking (
    id SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL,
    time_slot_id INTEGER NOT NULL,
    cancelled BOOLEAN NOT NULL, 
    created_at TIMESTAMP DEFAULT now()
);