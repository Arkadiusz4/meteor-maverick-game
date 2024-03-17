# Meteor Maverick Game

Meteor Maverick is an arcade-style space shooter game developed using the Ebitengine framework. In this game, players
control a spaceship tasked with defending Earth from incoming waves of meteoroids and enemy spacecraft.

## Features

- Intuitive Controls: Simple controls make it easy for players to maneuver their spaceship and engage in combat.
- Power-Ups: Collect power-ups scattered throughout the game to upgrade your weapons and enhance your abilities.
- Challenging Enemies: Battle against a variety of enemy spacecraft and dodge incoming meteoroids to survive each wave.
- Dynamic Levels: Encounter dynamically generated levels with increasing difficulty as you progress through the game.

## Installation

To run the game locally, follow these steps:

1. Clone the repository to your local machine:

```bash
git clone https://github.com/Arkadiusz4/meteor-maverick-game.git
```

2. Navigate to the project directory:

```bash
cd meteor-maverick-game
```

3. Run the main file:

```bash
go run cmd/main.go
```

## Controls

- **Arrow Keys**: Move the spaceship
- **Spacebar**: Fire primary weapon

## Testing

This project includes unit tests to ensure its functionality remains intact. The tests are located in the tests
directory and can be run using the following command:

```bash
cd tests
```

```bash
go test ./...
```

Additionally, continuous integration (**GitHub Actions**) is set up to run these tests automatically every day to
maintain the project's stability and reliability.

## Contributing

Contributions to Meteor Maverick are welcome! If you encounter any bugs or have suggestions for improvements, please
open an issue on the GitHub repository.

## Credits

Meteor Maverick was created by **Arkadiusz4**.

## License

This project is licensed under the **MIT License** - see the LICENSE file for details.
