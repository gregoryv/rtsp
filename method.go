package rtsp

//go:generate stringer -type Method -output methodlines.go
const (
	DESCRIBE      Method = iota // C->S             P,S       recommended
	ANNOUNCE                    // C->S, S->C       P,S       optional
	GET_PARAMETER               // C->S, S->C       P,S       optional
	OPTIONS                     // C->S, S->C       P,S       required, optional
	PAUSE                       // C->S             P,S       recommended
	PLAY                        // C->S             P,S       required
	RECORD                      // C->S             P,S       optional
	REDIRECT                    //       S->C       P,S       optional
	SETUP                       // C->S             S         required
	SET_PARAMETER               // C->S, S->C       P,S       optional
	TEARDOWN                    // C->S             P,S       required
)

type Method int
