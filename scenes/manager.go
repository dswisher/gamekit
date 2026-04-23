package scenes

// Ensure SceneManager implements SceneController interface.
var _ SceneController = (*SceneManager)(nil)

// SceneManager provides the ability to push new scenes and pop existing ones.
type SceneManager struct {
	scenes []Scene
}

// NewSceneManager creates a new scene manager and returns a pointer to it.
func NewSceneManager() *SceneManager {
	return &SceneManager{}
}

// Push adds a scene on top of the stack.
func (s *SceneManager) Push(sc Scene) {
	s.scenes = append(s.scenes, sc)
	sc.Enter()
}

// Pop removes the top scene.
func (s *SceneManager) Pop() {
	if len(s.scenes) == 0 {
		return
	}
	top := s.scenes[len(s.scenes)-1]
	top.Exit()
	s.scenes = s.scenes[:len(s.scenes)-1]
}

// Current returns the top scene without removing it.
func (s *SceneManager) Current() Scene {
	if len(s.scenes) == 0 {
		return nil
	}
	return s.scenes[len(s.scenes)-1]
}

// Replace pops current and pushes new.
func (s *SceneManager) Replace(sc Scene) {
	s.Pop()
	s.Push(sc)
}

// GetScenes returns a copy of all scenes in the stack.
func (s *SceneManager) GetScenes() []Scene {
	result := make([]Scene, len(s.scenes))
	copy(result, s.scenes)
	return result
}
