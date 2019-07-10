## Systemd formatter for logrus
Formats logrus entries into systemd style log lines.

    <3>an error happened

### Log level mapping
It maps:

    PanicLevel -> EMER(0)
    FatalLevel -> CRIT(2)
    ErrorLevel -> ERR(3)
    WarnLevel -> WARNING(4)
    InfoLevel -> INFO(6)
    DebugLevel -> DEBUG(7)
    TraceLevel -> DEBUG(7)

