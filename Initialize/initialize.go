package Initialize

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
)

func Initialize(serviceName string) {
	eg := errgroup.Group{}
	eg.Go(func() error {
		err := generateDockerfile()
		if err != nil {
			return err
		}
		fmt.Println("Generated Dockerfile")
		return nil
	})
	// ---
	eg.Go(func() error {
		err := generateDockerfileDev()
		if err != nil {
			return err
		}
		fmt.Println("Generated Dockerfile.dev")
		return nil
	})
	// ---
	eg.Go(func() error {
		err := generateDockerCompose()
		if err != nil {
			return err
		}
		fmt.Println("Generated docker-compose.yaml")
		return nil
	})
	// ---
	eg.Go(func() error {
		err := generateKubernetesManifest(serviceName)
		if err != nil {
			return err
		}
		fmt.Println("Generated Kubernetes Manifest")
		return nil
	})
	// ---
	eg.Go(func() error {
		err := initializeDirectory()
		if err != nil {
			return err
		}
		fmt.Println("Generated Directory Template")
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalln(err)
	}
}

func generateDockerfile() error {
	file, err := os.OpenFile("Dockerfile", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, Dockerfile)
	if err != nil {
		return err
	}
	return nil
}

func generateDockerfileDev() error {
	file, err := os.OpenFile("Dockerfile.dev", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, DockerfileDev)
	if err != nil {
		return err
	}
	return nil
}

func generateKubernetesManifest(serviceName string) error {
	err := os.Mkdir("k8s", 0777)
	if err != nil {
		return err
	}
	eg := errgroup.Group{}
	eg.Go(func() error {
		file, err := os.OpenFile("./k8s/" + serviceName + "_deployment.yaml", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = fmt.Fprintln(file, KubernetesDeploymentManifest)
		if err != nil {
			return err
		}
		return nil
	})
	// ---
	eg.Go(func() error {
		file, err := os.OpenFile("./k8s/" + serviceName + "_service.yaml", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = fmt.Fprintln(file, KubernetesServiceManifest)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func generateDockerCompose() error {
	file, err := os.OpenFile("docker-compose.yaml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, DockerCompose)
	if err != nil {
		return err
	}
	return nil
}

func initializeDirectory() error {
	eg := errgroup.Group{}
	eg.Go(func() error {
		err := os.Mkdir("application", 0777)
		if err != nil {
			return err
		}
		err = os.Mkdir("./application/controllers", 0777)
		if err != nil {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		err := os.Mkdir("domain", 0777)
		if err != nil {
			return err
		}
		err = os.Mkdir("./domain/services", 0777)
		if err != nil {
			return err
		}
		err = os.Mkdir("./domain/models", 0777)
		if err != nil {
			return err
		}
		return nil
	})
	eg.Go(func() error {
		err := os.Mkdir("infrastructure", 0777)
		if err != nil {
			return err
		}
		err = os.Mkdir("./infrastructure/persistence", 0777)
		if err != nil {
			return err
		}
		err = os.Mkdir("./infrastructure/persistence/repositories", 0777)
		if err != nil {
			return err
		}
		err = os.Mkdir("./infrastructure/persistence/dbmodels", 0777)
		if err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
