package consensus

import (
	"github.com/pactus-project/pactus/crypto/hash"
	"github.com/pactus-project/pactus/types/proposal"
	"github.com/pactus-project/pactus/types/vote"
)

type cpPreVoteState struct {
	*consensus
}

func (s *cpPreVoteState) enter() {
	s.decide()
}

func (s *cpPreVoteState) decide() {
	if s.cpRound == 0 {
		// broadcast the initial value
		prepares := s.log.PrepareVoteSet(s.round)
		preparesQH := prepares.QuorumHash()
		if preparesQH != nil {
			s.cpWeakValidity = preparesQH
			cert := s.makeCertificate(prepares.BlockVotes(*preparesQH))
			just := &vote.JustInitZero{
				QCert: cert,
			}
			s.signAddCPPreVote(*s.cpWeakValidity, s.cpRound, 0, just)
		} else {
			just := &vote.JustInitOne{}
			s.signAddCPPreVote(hash.UndefHash, s.cpRound, 1, just)
		}
	} else {
		cpMainVotes := s.log.CPMainVoteVoteSet(s.round)
		if cpMainVotes.HasAnyVoteFor(s.cpRound-1, vote.CPValueOne) {
			s.logger.Info("cp: one main-vote for one", "b", "1")

			vote1 := cpMainVotes.GetRandomVote(s.cpRound-1, vote.CPValueOne)
			just1 := &vote.JustPreVoteHard{
				QCert: vote1.CPJust().(*vote.JustMainVoteNoConflict).QCert,
			}
			s.signAddCPPreVote(hash.UndefHash, s.cpRound, vote.CPValueOne, just1)
		} else if cpMainVotes.HasAnyVoteFor(s.cpRound-1, vote.CPValueZero) {
			s.logger.Info("cp: one main-vote for zero", "b", "0")

			vote0 := cpMainVotes.GetRandomVote(s.cpRound-1, vote.CPValueZero)
			just0 := &vote.JustPreVoteHard{
				QCert: vote0.CPJust().(*vote.JustMainVoteNoConflict).QCert,
			}
			s.signAddCPPreVote(*s.cpWeakValidity, s.cpRound, vote.CPValueZero, just0)
		} else if cpMainVotes.HasAllVotesFor(s.cpRound-1, vote.CPValueAbstain) {
			s.logger.Info("cp: all main-votes are abstain", "b", "0 (biased)")

			votes := cpMainVotes.BinaryVotes(s.cpRound-1, vote.CPValueAbstain)
			cert := s.makeCertificate(votes)
			just := &vote.JustPreVoteSoft{
				QCert: cert,
			}
			s.signAddCPPreVote(*s.cpWeakValidity, s.cpRound, vote.CPValueZero, just)
		} else {
			s.logger.Panic("protocol violated. We have combination of votes for one and zero")
		}
	}

	s.enterNewState(s.cpMainVoteState)
}

func (s *cpPreVoteState) onAddVote(_ *vote.Vote) {
	panic("Unreachable")
}

func (s *cpPreVoteState) onSetProposal(_ *proposal.Proposal) {
	panic("Unreachable")
}

func (s *cpPreVoteState) onTimeout(_ *ticker) {
	panic("Unreachable")
}

func (s *cpPreVoteState) name() string {
	return "cp:pre-vote"
}