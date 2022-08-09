package demo;

import org.junit.Test;

class Pet {
    public void speak() {
        System.out.print("...");
    }

    public void speakTo(String name) {
        this.speak();
        System.out.println(name);
    }
}

class Dog extends Pet {
    @Override
    public void speak() {
        System.out.print("Wang!");
    }
}

public class InheritanceTest {
    @Test
    public void testSubClassAccess() {
        Pet aDog = new Dog();
        aDog.speak();
        aDog.speakTo("Chao");
    }

    private void makePetSpeak(Pet p) {
        p.speak();
        System.out.println("\nPet spoke.");
    }

    @Test
    public void testLSP() {
        Dog aDog = new Dog();
        makePetSpeak(aDog);
    }
}
