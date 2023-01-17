const displayContent= require('./load_reload_scripts');

describe("displayContent", () => {
    beforeEach(() => {
      global.fetch = jest.fn().mockImplementation(() =>
        Promise.resolve({
          json: () => Promise.resolve({ code:200,message: "Stack Content", content: ["1", "2", "3"] })
        })
      );
      global.resultDiv = { innerHTML: "" };
      jest.spyOn(global.console, "log");
    });
  
    afterEach(() => {
      global.fetch.mockClear();
      global.console.log.mockClear();
    });
  
    it("should make a fetch request to https:localhost:3000/display", async () => {
      await displayContent();
      expect(global.fetch).toHaveBeenCalledWith("https:localhost:3000/display");
    });
  
    it("should log the returned data", async () => {
      await displayContent();
      expect(global.console.log).toHaveBeenCalledWith({ code:200,message: "Stack Content", content: ["1", "2", "3"] });
    });
  
    it("should set resultDiv innerHTML with message and reversed content", async () => {
      await displayContent();
      expect(global.resultDiv.innerHTML).toBe("Stack Content:<p style='display:inline'> &nbsp3</p><p style='display:inline'> &nbsp2</p><p style='display:inline'> &nbsp1</p>");
    });
  });