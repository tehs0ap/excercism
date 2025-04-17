public static class PhoneNumber
{
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber)
    {
        bool isNewYork = false;
        bool isFake = false;
        string localNumber;
        
        var parts = phoneNumber.Split('-');

        isNewYork = parts[0] == "212";
        isFake = parts[1] == "555";
        localNumber = parts[2];
        
        return (isNewYork, isFake, localNumber);
    }

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) phoneNumberInfo)
    {
        return  phoneNumberInfo.IsFake;
    }
}
