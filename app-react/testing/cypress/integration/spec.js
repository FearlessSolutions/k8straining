it('detects Fearless on the page', () => {
    cy.visit('/');

    cy.contains('Fearless');
});

