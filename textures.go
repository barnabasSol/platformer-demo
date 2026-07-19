package main

func getTextures(state string) []string {

	if state == "idle" {
		return []string{
			"./res/fighter_sprites/fighter_Idle_0001.png",
			"./res/fighter_sprites/fighter_Idle_0002.png",
			"./res/fighter_sprites/fighter_Idle_0003.png",
			"./res/fighter_sprites/fighter_Idle_0004.png",
			"./res/fighter_sprites/fighter_Idle_0005.png",
			"./res/fighter_sprites/fighter_Idle_0006.png",
			"./res/fighter_sprites/fighter_Idle_0007.png",
			"./res/fighter_sprites/fighter_Idle_0008.png",
		}
	}

	if state == "walk" {
		return []string{
			"./res/fighter_sprites/fighter_walk_0009.png",
			"./res/fighter_sprites/fighter_walk_0010.png",
			"./res/fighter_sprites/fighter_walk_0011.png",
			"./res/fighter_sprites/fighter_walk_0012.png",
			"./res/fighter_sprites/fighter_walk_0013.png",
			"./res/fighter_sprites/fighter_walk_0014.png",
			"./res/fighter_sprites/fighter_walk_0015.png",
			"./res/fighter_sprites/fighter_walk_0016.png",
		}
	}

	if state == "run" {
		return []string{
			"./res/fighter_sprites/fighter_run_0017.png",
			"./res/fighter_sprites/fighter_run_0018.png",
			"./res/fighter_sprites/fighter_run_0019.png",
			"./res/fighter_sprites/fighter_run_0020.png",
			"./res/fighter_sprites/fighter_run_0021.png",
			"./res/fighter_sprites/fighter_run_0022.png",
			"./res/fighter_sprites/fighter_run_0023.png",
			"./res/fighter_sprites/fighter_run_0024.png",
		}
	}

	if state == "jump" {
		return []string{
			"./res/fighter_sprites/fighter_jump_0043.png",
			"./res/fighter_sprites/fighter_jump_0044.png",
			"./res/fighter_sprites/fighter_jump_0045.png",
			"./res/fighter_sprites/fighter_jump_0046.png",
			"./res/fighter_sprites/fighter_jump_0047.png",
		}

	}

	if state == "dash" {
		return []string{
			"./res/fighter_sprites/fighter_dash_0033.png",
			"./res/fighter_sprites/fighter_dash_0034.png",
			"./res/fighter_sprites/fighter_dash_0035.png",
			"./res/fighter_sprites/fighter_dash_0036.png",
			"./res/fighter_sprites/fighter_dash_0037.png",
			"./res/fighter_sprites/fighter_dash_0038.png",
		}
	}
	return []string{}
}
