using GTANetworkAPI;

public class Events : Script
{
    [ServerEvent(Event.ResourceStart)]
    public void OnResourceStart()
    {
        NAPI.Util.ConsoleOutput("OnResourceStart");
    }
}