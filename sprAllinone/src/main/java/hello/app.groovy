import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class ThisWillActuallyRun {

    @RequestMapping("/gr")
    String home() {
        return "Hello world for groovy!";
    }
}


