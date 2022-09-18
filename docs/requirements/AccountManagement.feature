Feature: Account Management

    Scenario: Register Bank Account
        Given I am a user
        And I opened finkita application 
        And I navigated to account page
        Then I should see a button that says "Add Account"


