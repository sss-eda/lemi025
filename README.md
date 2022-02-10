# LEMI-025

This repository should contain everything that is specific to the LEMI-025 instrument. That includes the frontend - components that will provide the user with client interface.

Control commands sent by the user from the user interface should be considered as queries and not as commands. The read model these queries are aimed at simply uses the serial interface (or some other representation) as the gateway (secondary adapter).

The device send the results from the commands executed over serial port as commands to the write model (aggregate).