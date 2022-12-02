package main

/*
*****************************************************************

	@author: Adenugba Adeoluwa
	Date: 2nd  December 2022
	Purpose: To run an election which calculates the vote and prints out the winner
********************************************************************
*/

import (
	"fmt"
	"os"
)

type candidate struct {
	name          string
	vote_total    int
	precint_votes []votes
}

type votes struct {
	precint byte
	tally   int
}

var winner = []string{}
var toptwo = map[string]int{}
var overall_votes = 0

var runoff = map[string]int{}
func main() {
	// initializing and declaring candidates information
	candidate1 := candidate{
		name:       "candidate1",
		vote_total: 0,
		precint_votes: []votes{
			votes{
				precint: 'A',
				tally:   192,
			},
			votes{
				precint: 'B',
				tally:   147,
			},
			votes{
				precint: 'C',
				tally:   186,
			},
			votes{
				precint: 'D',
				tally:   11,
			},
			votes{
				precint: 'E',
				tally:   26,
			},
		},
	}

	candidate2 := candidate{
		name:       "candidate2",
		vote_total: 0,
		precint_votes: []votes{
			votes{
				precint: 'A',
				tally:   48,
			},
			votes{
				precint: 'B',
				tally:   90,
			},
			votes{
				precint: 'C',
				tally:   12,
			},
			votes{
				precint: 'D',
				tally:   21,
			},
			votes{
				precint: 'E',
				tally:   13,
			},
		},
	}

	candidate3 := candidate{
		name:       "candidate3",
		vote_total: 0,
		precint_votes: []votes{
			votes{
				precint: 'A',
				tally:   20,
			},
			votes{
				precint: 'B',
				tally:   31,
			},
			votes{
				precint: 'C',
				tally:   12,
			},
			votes{
				precint: 'D',
				tally:   40,
			},
			votes{
				precint: 'E',
				tally:   382,
			},
		},
	}

	candidate4 := candidate{
		name:       "candidate4",
		vote_total: 0,
		precint_votes: []votes{
			votes{
				precint: 'A',
				tally:   37,
			},
			votes{
				precint: 'B',
				tally:   21,
			},
			votes{
				precint: 'C',
				tally:   38,
			},
			votes{
				precint: 'D',
				tally:   39,
			},
			votes{
				precint: 'E',
				tally:   29,
			},
		},
	}
	fmt.Println("WELCOME TO THE ELECTION")

	for {

		fmt.Println("choose 1 option")
		fmt.Print(
			"[1] Show Precint Votes", "\n",
			"[2] Add Votes", "\n",
			"[3] Show Winner", "\n",
			"[4] Exit Program", "\n",
		)
		var value int
		fmt.Scan(&value)
		if value == 1 {
			candidate1.display_candidates_votes()
			candidate2.display_candidates_votes()
			candidate3.display_candidates_votes()
			candidate4.display_candidates_votes()

		} else if value == 2 {
			candidate1.add_votes()
			candidate2.add_votes()
			candidate3.add_votes()
			candidate4.add_votes()
		
			
			fmt.Println("VOTES ADDED SUCCESSFULLY")
		} else if value == 3 {
			candidate1.find_winner()
			candidate2.find_winner()
			candidate3.find_winner()
			candidate4.find_winner()

			if len(winner) == 0 {
				for i:=0;i<2;i++{
					tie()
					
				}
				fmt.Println("THE WINNERS ")
				for name,score:=range toptwo{
					
					fmt.Printf("%v: %v\n",name, score)
				}
				
			} else if len(winner) == 1 {
				fmt.Println(winner)
			}

		} else if value == 4 {
			os.Exit(1)
		}

	}

}

func (person *candidate) display_candidates_votes() {
	fmt.Print(person.name)
	for i := 0; i < 5; i++ {
		fmt.Printf(" %-3d ", person.precint_votes[i].tally)

	}
	fmt.Println("")
}

func (person *candidate) add_votes()  {
	
	for i := 0; i < 5; i++ {
		person.vote_total += person.precint_votes[i].tally
	}
	overall_votes += person.vote_total
	runoff[string(person.name)]=person.vote_total

}


func (person *candidate) find_winner() {

	if person.vote_total > overall_votes/2 {
		winner = append(winner, person.name)
	}
}
func tie(){
	
	winners_name:= ""
	winners_vote:= 0
	for name,votes:= range runoff{
		
		
		if votes> winners_vote{
			winners_vote = votes
			winners_name = name
		}
	}
	toptwo[string(winners_name)]=winners_vote
	delete(runoff,winners_name)
	winners_vote=0
	winners_name=""
	

}
