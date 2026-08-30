package main

// player
func getTextures(state string) []string {
	if state == StateIdle {
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

	if state == StateWalk {
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

	if state == StateRun {
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

	if state == StateJump {
		return []string{
			"./res/fighter_sprites/fighter_jump_0043.png",
			"./res/fighter_sprites/fighter_jump_0044.png",
			"./res/fighter_sprites/fighter_jump_0045.png",
			"./res/fighter_sprites/fighter_jump_0046.png",
			"./res/fighter_sprites/fighter_jump_0047.png",
			"./res/fighter_sprites/fighter_jump_0043.png",
		}
	}

	if state == StateDash {
		return []string{
			"./res/fighter_sprites/fighter_dash_0033.png",
			"./res/fighter_sprites/fighter_dash_0034.png",
			"./res/fighter_sprites/fighter_dash_0035.png",
			"./res/fighter_sprites/fighter_dash_0036.png",
			"./res/fighter_sprites/fighter_dash_0037.png",
			"./res/fighter_sprites/fighter_dash_0038.png",
		}
	}
	if state == StateSlide {
		return []string{
			"./res/fighter_sprites/fighter_slide_0025.png",
			"./res/fighter_sprites/fighter_slide_0026.png",
			"./res/fighter_sprites/fighter_slide_0027.png",
			"./res/fighter_sprites/fighter_slide_0028.png",
			"./res/fighter_sprites/fighter_slide_0029.png",
			"./res/fighter_sprites/fighter_slide_0030.png",
			"./res/fighter_sprites/fighter_slide_0031.png",
			"./res/fighter_sprites/fighter_slide_0032.png",
		}
	}
	if state == StateCombo1 {
		return []string{
			"./res/fighter_sprites/fighter_combo_0064.png",
			"./res/fighter_sprites/fighter_combo_0065.png",
			"./res/fighter_sprites/fighter_combo_0066.png",
			"./res/fighter_sprites/fighter_combo_0067.png",
			"./res/fighter_sprites/fighter_combo_0068.png",
		}
	}
	if state == StateCombo2 {
		return []string{
			"./res/fighter_sprites/fighter_combo_0069.png",
			"./res/fighter_sprites/fighter_combo_0070.png",
			"./res/fighter_sprites/fighter_combo_0071.png",
			"./res/fighter_sprites/fighter_combo_0072.png",
			"./res/fighter_sprites/fighter_combo_0073.png",
			"./res/fighter_sprites/fighter_combo_0074.png",
		}
	}
	if state == StateCombo3 {
		return []string{
			"./res/fighter_sprites/fighter_combo_0075.png",
			"./res/fighter_sprites/fighter_combo_0076.png",
			"./res/fighter_sprites/fighter_combo_0077.png",
			"./res/fighter_sprites/fighter_combo_0078.png",
			"./res/fighter_sprites/fighter_combo_0079.png",
			"./res/fighter_sprites/fighter_combo_0080.png",
			"./res/fighter_sprites/fighter_combo_0081.png",
			"./res/fighter_sprites/fighter_combo_0082.png",
		}
	}
	return []string{}
}
