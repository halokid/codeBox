package org.halokid.transactional;

import org.springframework.data.jpa.repository.JpaRepository;

public interface SpringBootUserRepository extends JpaRepository<JpaSpringBootUser, Integer> {
}


