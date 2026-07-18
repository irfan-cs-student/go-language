package main

import "fmt"

func test() {
	x := 100
	fmt.Println(x)
}

func main() {

	fmt.Println("========================================")
	fmt.Println("        GO TOUGH SCOPE CHALLENGES")
	fmt.Println("========================================")

	// =========================================
	// Challenge 11 - Triple Shadowing
	// =========================================

	fmt.Println("\nChallenge 11")

	{
		a := 1

		{
			a := a + 1

			{
				a := a + 1

				{
					a := a + 1
					fmt.Println(a) //4
				}

				fmt.Println(a) //3
			}

			fmt.Println(a) //2
		}

		fmt.Println(a) //1
	}

	// Predict:
	//

	// =========================================
	// Challenge 12 - Which x changes?
	// =========================================

	fmt.Println("\nChallenge 12")

	{
		x := 100

		{
			x := 200

			{
				x = 300
			}

			fmt.Println(x) //300
		}

		fmt.Println(x) // 100
	}

	// Predict:
	//

	// =========================================
	// Challenge 13 - Mixed Assignment
	// =========================================

	fmt.Println("\nChallenge 13")

	{
		a := 5

		{
			a = 10

			{
				a := 20

				{
					a = 30
				}

				fmt.Println(a) //30
			}

			fmt.Println(a) //10
		}

		fmt.Println(a) //10
	}

	// Predict:
	//

	// =========================================
	// Challenge 14 - Same Name Everywhere
	// =========================================

	fmt.Println("\nChallenge 14")

	{
		x := 1

		{
			x := 2

			{
				x := 3

				{
					x := 4

					{
						fmt.Println(x) //44
					}
				}

				fmt.Println(x) //3
			}

			fmt.Println(x) //2
		}

		fmt.Println(x) //1
	}

	// Predict:
	//

	// =========================================
	// Challenge 15 - Sneaky
	// =========================================

	fmt.Println("\nChallenge 15")

	/*
		{
			a := 10

			{
				fmt.Println(a) //10

				a := 20

				fmt.Println(a) //20
			}

			fmt.Println(a) //10
		}
	*/

	// Compile?
	// Output?

	// =========================================
	// Challenge 16 - if Scope
	// =========================================

	fmt.Println("\nChallenge 16")

	{
		x := 5

		if x := x + 10; x > 10 {
			fmt.Println(x)
		}

		fmt.Println(x) //5
	}

	// Predict:
	//

	// =========================================
	// Challenge 17 - Loop Scope
	// =========================================

	fmt.Println("\nChallenge 17")

	{
		i := 100

		for i := 1; i <= 3; i++ {
			fmt.Println(i)
		}

		fmt.Println(i) //100
	}

	// Predict:
	//

	// =========================================
	// Challenge 18 - Nested Loops
	// =========================================

	fmt.Println("\nChallenge 18")

	{
		x := 1

		for i := 1; i <= 2; i++ {

			x := x + i

			for j := 1; j <= 2; j++ {
				x = x + j
			}

			fmt.Println(x)
		}

		fmt.Println(x) //1
	}

	// Predict:
	//

	// =========================================
	// Challenge 19 - Function Scope
	// =========================================

	fmt.Println("\nChallenge 19")

	{
		x := 10

		test()

		fmt.Println(x) //10
	}

	// Predict:
	//

	// =========================================
	// Challenge 20 - Boss Fight
	// =========================================

	fmt.Println("\nChallenge 20")

	{
		a := 1

		{
			a = 2

			{
				a := a + 2 //4

				{
					a = a + 2 //a=6

					{
						a := a + 2 //8

						{
							a = a + 2      //10
							fmt.Println(a) //10
						}

						fmt.Println(a) //10
					}

					fmt.Println(a) //6
				}

				fmt.Println(a) //6
			}

			fmt.Println(a) //2
		}

		fmt.Println(a) //2
	}

	// Predict:
	//

	// =========================================
	// Ultimate Challenge
	// =========================================

	fmt.Println("\nUltimate Challenge")

	{
		x := 1

		{
			x = 2

			{
				x := x + 3 //5

				{
					x := x + 4 //9

					{
						x = x + 5 //14
					}

					fmt.Println("A", x) //a,14
				}

				fmt.Println("B", x) //5

				{
					x = x + 6 //11
				}

				fmt.Println("C", x) //11
			}

			fmt.Println("D", x) //2
		}

		fmt.Println("E", x) //2
	}

	// Predict:
	//

	// =========================================
	// FINAL INTERVIEW QUESTION
	// =========================================

	fmt.Println("\nFinal Interview Question")

	/*
		a := 10

		{
			a := 20

			{
				a = 30
			}
		}

		Questions:
		1. How many variables named 'a' exist? //2
		2. Which one becomes 30?//second 1
		3. When is it destroyed?//after moving out of parenthesis
		4. Final value of outer a?//10
	*/

}
