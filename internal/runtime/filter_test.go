package runtime

import (
	"testing"
)

func TestJsFilter(t *testing.T) {
	// 	workerID := "yo"
	// 	now := time.Now()
	// 	participants := []*ent.Participant{
	// 		{
	// 			ID:            "123",
	// 			CreatedAt:     now,
	// 			UpdatedAt:     now,
	// 			MturkWorkerID: &workerID,
	// 			Edges: ent.ParticipantEdges{
	// 				Data: []*ent.Datum{
	// 					{
	// 						ID:        "hah",
	// 						CreatedAt: time.Now(),
	// 						UpdatedAt: time.Now(),
	// 						Key:       "everything",
	// 						Val:       "42",
	// 						Index:     0,
	// 						Current:   true,
	// 						Version:   0,
	// 					},
	// 				},
	// 			},
	// 		},
	// 		{
	// 			ID:            "456",
	// 			CreatedAt:     now,
	// 			UpdatedAt:     now,
	// 			MturkWorkerID: &workerID,
	// 			Edges: ent.ParticipantEdges{
	// 				Data: []*ent.Datum{
	// 					{
	// 						ID:        "hah",
	// 						CreatedAt: time.Now(),
	// 						UpdatedAt: time.Now(),
	// 						Key:       "everything",
	// 						Val:       "20",
	// 						Index:     0,
	// 						Current:   true,
	// 						Version:   0,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	}

	// 	jsfilter(context.Background(), participants, `function (participants, arg1, arg2) {
	// 		participants[0].set("hah", "789")
	// 		console.log(participants[0].get("hah"))
	// 		participants[1].set("everything", 951)
	// 		console.log(participants[1].get("everything"))
	// 		return participants.filter(p => p.get("everything") < 300);
	// }`)
}
