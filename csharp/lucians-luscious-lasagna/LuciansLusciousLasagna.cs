class Lasagna
{
    private readonly int _expectedMinutesInOven = 40;
    private readonly int _layerMinutesPrepTime = 2;
    public int ExpectedMinutesInOven()
    {
        return _expectedMinutesInOven;
    }
    
    public int RemainingMinutesInOven(int minutesInOven)
    {
        return ExpectedMinutesInOven() - minutesInOven;
    }
    
    public int PreparationTimeInMinutes(int layers)
    {
        return layers * _layerMinutesPrepTime;
    }

    public int ElapsedTimeInMinutes(int layers, int  minutesInOven)
    {
        return PreparationTimeInMinutes(layers) + minutesInOven;
    }
}
