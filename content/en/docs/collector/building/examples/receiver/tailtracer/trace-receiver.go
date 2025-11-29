package tailtracer

import (
    "context"
    "go.opentelemetry.io/collector/component"
)

# implement Start() & Shutdown()
type tailtracerReceiver struct{

}

func (tailtracerRcvr *tailtracerReceiver) Start(ctx context.Context, host component.Host) error {
    return nil
}

func (tailtracerRcvr *tailtracerReceiver) Shutdown(ctx context.Context) error {
    return nil
}
